package tasks

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"
)

type Request struct {
	Payload string
}

type Client interface {
	SendRequest(ctx context.Context, request Request) error
}

type client struct {
}

type RateLimiter struct {
	ticker *time.Ticker
	rps    int
}

func NewRateLimiter(rps int) *RateLimiter {
	ticker := time.NewTicker(time.Second / time.Duration(rps))
	return &RateLimiter{
		rps:    rps,
		ticker: ticker,
	}
}

func (rl *RateLimiter) allow() {
	<-rl.ticker.C
}

func (c client) SendRequest(ctx context.Context, request Request) error {
	fmt.Println("send request: ", request.Payload)
	return nil
}

func RateLimiterTask() {
	c := client{}
	ctx := context.Background()
	var requests []Request
	for i := 0; i < 500; i++ {
		request := Request{
			strconv.Itoa(i),
		}
		requests = append(requests, request)
	}

	makeBatchApiCalls(ctx, c, requests)

}

func makeBatchApiCalls(ctx context.Context, c Client, requests []Request) {
	var wg sync.WaitGroup
	rl := NewRateLimiter(100)
	t := time.Now()
	for _, r := range requests {
		r := r
		wg.Add(1)
		go func() {
			defer func() {
				wg.Done()
			}()
			rl.allow()
			err := c.SendRequest(ctx, r)
			if err != nil {
				log.Println("Ошибка: ", err)
			}
		}()
	}

	wg.Wait()

	log.Println(time.Since(t))
}
