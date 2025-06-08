package main

import (
	"strconv"
)

func main() {
	compress([]byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'd'})
}

func compress(chars []byte) int {
	start := 0
	count := 1
	pos := 0

	for i := 1; i < len(chars); i++ {
		if chars[i] != chars[i-1] {
			chars[pos] = chars[i-1]
			pos++
			if count > 1 {
				res := []byte(strconv.Itoa(i - start))
				for _, v := range res {
					chars[pos] = v
					pos++
				}
			}
			count = 1
		} else {
			count++
		}
	}

	return len(chars[:pos])
}
