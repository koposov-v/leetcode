package main

import (
	"fmt"
	"sync"
)

type WordCounter struct {
	counts   map[string]int
	lastWord string
	limit    int
}

func NewWorkCounter(limit int) *WordCounter {
	return &WordCounter{
		limit:  limit,
		counts: make(map[string]int, limit),
	}
}

func (wc *WordCounter) CountWord(word string) {
	sync.Map{}
	if _, ok := wc.counts[word]; !ok {
		wc.lastWord = word
	}

	wc.counts[word]++

	if len(wc.counts) > wc.limit {
		delete(wc.counts, wc.lastWord)
	}
}

func main() {
	wc := NewWorkCounter(3)

	words := []string{"apple", "banana", "apple", "kiwi", "orange", "banana"}

	for _, word := range words {
		wc.CountWord(word)
	}

	fmt.Println("Слова:", wc.counts)
	fmt.Println("Кол-во слов:", wc.limit)
}
