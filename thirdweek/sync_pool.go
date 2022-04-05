package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func text_pool() {
	pool := sync.Pool{
		New: func() interface{} {
			return "123456"
		},
	}
	pool.Put("123")
	pool.Put("345")
	pool.Put("789")
	pool.Put("101")
	pool.Put("000")
	runtime.GC()
	time.Sleep(time.Second * 1)
	fmt.Println(pool.Get().(string))
	fmt.Println(pool.Get().(string))
	time.Sleep(time.Second * 1)
	runtime.GC()
	pool.Put("101")
	fmt.Println(pool.Get().(string))
	fmt.Println(pool.Get().(string))
	runtime.GC()
	time.Sleep(time.Second * 1)
	fmt.Println(pool.Get().(string))
}
