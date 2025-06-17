package main

import (
	"errors"
	"strings"
	"time"
)

func main() {
	f := 3
	limiter := NewLimiter(f, 10*time.Second)
	yesterday := time.Now().Truncate(24 * time.Hour)
	now := time.Now()
	limiter.Log("1", yesterday)
	limiter.Log("1", now)
	limiter.Log("1", now.Add(time.Second))
	limiter.Log("1", now.Add(time.Second))
	limiter.Log("1", now.Add(time.Second))
	limiter.Check("1")
}

type Limiter struct {
	limit    int
	interval time.Duration
	requests map[string][]time.Time
}

func NewLimiter(limit int, interval time.Duration) *Limiter {
	return &Limiter{
		limit:    limit,
		interval: interval,
		requests: make(map[string][]time.Time),
	}
}

func (l *Limiter) Log(ip string, date time.Time) {
	if times, ok := l.requests[ip]; ok {
		l.requests[ip] = append(times, date)
		return
	}

	l.requests[ip] = []time.Time{date}
}

func (l *Limiter) Check(ip string) error {
	times, ok := l.requests[ip]
	if !ok {
		return nil
	}
	now := time.Now()
	start := 0

	for _, t := range times {
		if now.Sub(t) <= l.interval {
			times[start] = t
			start++
		}
	}

	l.requests[ip] = times[:start+1]
	if len(l.requests[ip]) > l.limit {
		return errors.New("too Many Requests")
	}

	return nil
}

func mergeAlternately(word1 string, word2 string) string {
	builder := strings.Builder{}
	i := 0
	j := 0
	k := 0
	for i < len(word1) && j < len(word2) {
		isEven := (k+1)%2 == 0
		if len(word1) > i && !isEven {
			builder.WriteByte(word1[i])
		}
		if len(word2) > j && isEven {
			builder.WriteByte(word2[j])
		}
		k++
	}
	return builder.String()
}

func reverseVowels(s string) string {
	ss := strings.Fields(s)
	left := 0
	right := len(ss) - 1
	hashMap := map[string]struct{}{
		"a": struct{}{},
		"e": struct{}{},
		"i": struct{}{},
		"o": struct{}{},
		"u": struct{}{},
		"A": struct{}{},
		"E": struct{}{},
		"I": struct{}{},
		"O": struct{}{},
		"U": struct{}{},
	}

	for left < right {
		_, existsLeft := hashMap[ss[left]]
		_, existsRight := hashMap[ss[right]]
		if existsLeft && existsRight {
			ss[left], ss[right] = ss[right], ss[left]
		}
	}

	return strings.Join(ss, "")
}
