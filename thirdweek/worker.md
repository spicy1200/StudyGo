package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/go-kratos/kratos/pkg/sync/errgroup"
)

func main() {
	// 获取goroutin 上下文
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	groupContext := errgroup.WithContext(ctx)
	srv := &http.Server{Addr: ":8080"}
	groupContext.Go(func(ctx context.Context) error {
		return serverStart(srv)
	})
	groupContext.Go(func(ctx context.Context) error {
		<-ctx.Done()
		return srv.Shutdown(ctx)
	})
	channel := make(chan os.Signal, 1)

	signal.Notify(channel)

	groupContext.Go(func(ctx context.Context) error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-channel:
				cancel()
			}
		}
	})
	if err := groupContext.Wait(); err != nil {
		fmt.Println("group error", err)
	}
	fmt.Println("all group done!")

}

func serverStart(srv *http.Server) error {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("return host"))
	})
	err := srv.ListenAndServe()
	return err
}
