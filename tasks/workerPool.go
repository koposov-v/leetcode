package tasks

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	numJobs := 5
	jobs := make(chan int)
	results := make(chan int)
	wg := sync.WaitGroup{}

	for i := range numJobs {
		wg.Add(1)
		go worker(ctx, i+1, jobs, results, &wg)
	}

	go func() {
		for i := range 100 {
			select {
			case <-ctx.Done():
				break
			case jobs <- i + 1:
			}
		}

		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	for r := range results {
		fmt.Println("Result:", r)
	}

}

func worker(ctx context.Context, id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	resultCh := make(chan int)
	defer wg.Done()

	go func() {
		for j := range jobs {
			time.Sleep(time.Duration(rand.Intn(5+1)) * time.Second)
			resultCh <- j * 2
			fmt.Println("Worker:", id)
		}

		close(resultCh)
	}()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("finish worker by context, id: ", id)
			return
		case v, ok := <-resultCh:
			if !ok {
				return
			}
			results <- v
		}
	}
}
