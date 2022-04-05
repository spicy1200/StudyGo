package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-kratos/kratos/pkg/sync/errgroup"
)

func main() {
	var g errgroup.Group
	g.Go(func(ctx context.Context) error {
		time.Sleep(time.Second * 2)
		fmt.Println("exec #1")
		return nil
	})
	g.Go(func(ctx context.Context) error {
		time.Sleep(time.Second * 4)
		fmt.Println("exec # 2")
		return errors.New("failed to exec #2")
	})
	g.Go(func(ctx context.Context) error {
		time.Sleep(15 * time.Second)
		fmt.Println("exec #3")
		return nil
	})
	if err := g.Wait(); err == nil {
		fmt.Println(" successfully exec all")
	} else {
		fmt.Println("failed:", err)
	}
}
