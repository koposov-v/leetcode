package tasks

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu    *sync.Mutex
	value int
}

func NewCounter() Counter {
	mu := sync.Mutex{}
	return Counter{
		mu: &mu,
	}
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value += 1
}

func (c *Counter) Value() int {
	return c.value
}

func MutexTask() {

	conuter := NewCounter()
	var wg sync.WaitGroup
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			conuter.Increment()
		}()
	}
	wg.Wait()

	fmt.Printf("Counter value %d", conuter.Value())
}
