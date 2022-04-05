package main

import (
	"context"
	"fmt"
	"time"
)

type Tracker struct {
	ch   chan string
	stop chan struct{}
}

func NewTracker() *Tracker {
	return &Tracker{
		ch: make(chan string, 10),
	}
}
func (t *Tracker) Event(ctx context.Context, data string) error {
	select {

	case t.ch <- data:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
func (t *Tracker) Run() {
	for data := range t.ch {
		time.Sleep(1 * time.Second)
		fmt.Println(data)

	}
	t.stop <- struct{}{}
}
func (t *Tracker) Shutdown(ctx context.Context) {
	close(t.ch)
	select {
	case <-t.stop:
	case <-ctx.Done():
	}
}
func tracker() {
	tr := NewTracker()
	go tr.Run()
	_ = tr.Event(context.Background(), "text1")
	_ = tr.Event(context.Background(), "text2")
	_ = tr.Event(context.Background(), "text3")
	_ = tr.Event(context.Background(), "text4")
	time.Sleep(3 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	defer cancel()
	tr.Shutdown(ctx)
}
