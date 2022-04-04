package extend

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Config struct {
	a []int
}

func (c *Config) T() {
	fmt.Println("%v \n", c)
}

func atomicLock() {
	var v atomic.Value
	v.Store(&Config{})

	go func() {
		i := 0
		for {
			i++
			cfg := &Config{a: []int{i, i + 1, i + 2, i + 3, i + 4, i + 5}}
			v.Store(cfg)
		}
	}()
	var wg sync.WaitGroup
	for n := 0; n < 4; n++ {
		wg.Add(1)
		go func() {
			for n := 0; n < 100; n++ {
				cfg := v.Load().(*Config) // 断言 config 类型
				cfg.T()
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
