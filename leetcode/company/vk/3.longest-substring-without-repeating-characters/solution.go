package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	res := 0
    lastIndex := make([]int, 128)
    start := 0

    for end := 0; end < len(s); end++ {
        curChar := s[end]
        if lastIndex[curChar] > start {
            start = lastIndex[curChar]
        }
        if end - start + 1 > res {
            res = end - start + 1
        }

        lastIndex[curChar] = end + 1
    }

	return res
}

func main() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s))
}
