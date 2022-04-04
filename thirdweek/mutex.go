package main

import (
	"fmt"
	"sync"
)

type Config struct {
	a []int
}

func (c *Config) T() {
	fmt.Println(" %v \n", c)
}

//  读多写少就用读写锁  读特别特别多 可以使用 atomic
func setMutex() {
	var rwl sync.RWMutex
	var cfg *Config
	go func() {
		i := 0
		for {
			i++
			rwl.Lock()
			cfg = &Config{a: []int{i, i + 1, i + 2, i + 3, i + 4, i + 5}}
			rwl.Unlock()
		}
	}()
	var wg sync.WaitGroup
	for n := 0; n < 4; n++ {
		wg.Add(1)
		go func() {
			for n := 0; n < 6; n++ {
				rwl.Lock()
				cfg.T()
				rwl.Unlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
