package main

import (
	"context"
	"net/http"
	"sync"
)

func server(addr string, handler http.Handler, stop <-chan struct{}) error {
	s := http.Server{
		Addr:    addr,
		Handler: handler,
	}
	go func() {
		<-stop
		s.Shutdown(context.Background())
	}()

	return s.ListenAndServe()
}

func stopServer() {
	var wg sync.WaitGroup
	done := make(chan error, 2)
	stop := make(chan struct{})
	hf := http.NotFoundHandler()
	go func() {
		wg.Add(1)
		done <- server("127.0.0.1:8080", hf, stop)
	}()
	go func() {
		done <- server("127.0.0.1:8081", hf, stop)
	}()
	wg.Add(1)
	wg.Wait()
}
