package main

import (
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	builder := strings.Builder{}
	for _, v := range strings.ToLower(s) {
		if unicode.IsDigit(v) || unicode.IsLetter(v) {
			builder.WriteRune(v)
		}
	}

	cleanStr := builder.String()

	for i := 0; i < len(cleanStr)/2; i++ {
		if cleanStr[i] != cleanStr[len(cleanStr)-1-i] {
			return false
		}
	}
	return true
}
