package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

const goRun = 5

func main() {
	ctx := context.Background()
	_ = ctx

	ch := make(chan struct{}, goRun)

	wg := sync.WaitGroup{}

	for i := range 100 {
		wg.Add(1)
		ch <- struct{}{}
		go func() {
			defer wg.Done()
			defer func() { <-ch }()
			time.Sleep(time.Second * 1)
			fmt.Println(i)
		}()
	}

	wg.Wait()
	close(ch)
}
