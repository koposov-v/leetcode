package tasks

import (
	"fmt"
	"sync"
)

type Data struct {
	Id   int
	Name string
}

type Queue struct {
	chData chan Data
}

func NewQueue() *Queue {
	// Инициализация очереди и канала
	chData := make(chan Data)
	return &Queue{
		chData: chData,
	}
}

// Реализуйте метод добавления данных продюсерами

func (q *Queue) Produce(data Data) {
	q.chData <- data
}

func (q *Queue) CloseChanel() {
	close(q.chData)
}

func (q *Queue) Consume(wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range q.chData {
		fmt.Println(data)
	}
}

func ConsumeProduceTask() {
	var wg sync.WaitGroup
	queue := NewQueue()
	data := Data{
		Id:   123,
		Name: "Slava",
	}

	wg.Add(5)
	for i := 0; i < 5; i++ {
		go queue.Consume(&wg)
	}

	go func() {
		for i := 0; i < 10; i++ {
			queue.Produce(data)
		}
		queue.CloseChanel()
	}()

	wg.Wait()
}
