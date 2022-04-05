package main

import (
	"context"
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/go-kratos/kratos/pkg/sync/errgroup"
)

type result struct {
	path string
	sum  [md5.Size]byte
}

func errogroup_context() {
	// 读取文件信息
	data, err := MD5All(context.Background(), "./secondweek")
	if err != nil {
		log.Fatal(err)
	}
	for k, sum := range data {
		fmt.Println("%s:\t%x \n", k, sum)
	}

}
func MD5All(ctx context.Context, root string) (map[string][md5.Size]byte, error) {
	g := errgroup.WithContext(ctx)
	// 穿件文件路径通道
	paths := make(chan string)
	// 遍历文件，将文件路径放到Paths
	g.Go(func(ctx context.Context) error {
		defer close(paths)
		return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.Mode().IsRegular() {
				return nil
			}
			select {
			case paths <- path:
			case <-ctx.Done():
				return ctx.Err()
			}
			return nil
		})
	})
	// 20 个 goroutine计算md5 从paths获取文件路径
	c := make(chan result)
	const numDigesters = 20
	for i := 0; i < numDigesters; i++ {
		g.Go(func(ctx context.Context) error {
			for path := range paths {
				data, err := ioutil.ReadFile(path)
				if err != nil {
					return err
				}
				select {
				case c <- result{path, md5.Sum(data)}:
				case <-ctx.Done():
					return ctx.Err()
				}
			}
			return nil
		})
	}
	go func() {
		g.Wait() // 等待执行完成
		close(c)
	}()
	m := make(map[string][md5.Size]byte)
	for r := range c {
		m[r.path] = r.sum
	}
	if err := g.Wait(); err != nil {
		return nil, err
	}
	return m, nil
}
