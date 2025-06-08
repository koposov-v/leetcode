package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	chanles := make([]chan int64, 10)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(rand.Intn(4))*time.Second)
	defer cancel()

	for i := range chanles {
		ch := make(chan int64)
		go func(ch chan int64) {
			defer close(ch)

			select {
			case <-ctx.Done():
				return
			case ch <- int64(i):
			}
		}(ch)
		chanles[i] = ch
	}

	for ch := range merge(ctx, chanles...) {
		fmt.Println(ch)
	}
}

func merge(ctx context.Context, channels ...chan int64) <-chan int64 {
	result := make(chan int64)

	wg := sync.WaitGroup{}

	for i, ch := range channels {
		wg.Add(1)
		go func() {
			localCh := make(chan int64)
			defer wg.Done()

			go func() {
				for v := range ch {
					time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
					localCh <- v
				}
			}()

			for {
				select {
				case <-ctx.Done():
					fmt.Println("Закрыли горутину:", i)
					return
				case v, ok := <-localCh:
					if !ok {
						return
					}
					result <- v
				}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	return result
}
