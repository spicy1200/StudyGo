package main

import (
	"errors"
	"fmt"
	"time"
)

func search() (string, error) {
	time.Sleep(200 * time.Millisecond)
	return "some value", nil
}

type result struct {
	record string
	err    error
}

func main() {
	ch := make(chan int)
	go func() {
		record, err := search()
		ch <- result{record, err}
	}()
	select {
	case <-ctx.Done():
		return errors.New("search canceld")
	case request := <-ch:
		if request.err := nil{
			return result.err
		}
		fmt.Println("received:", request.record)
		return nil
	}
}
