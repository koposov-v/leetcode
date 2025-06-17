package main

import (
	"fmt"
)

func main() {
	fmt.Println("'a' ->", int('a')) // 97
	fmt.Println("'2' ->", int('2')) // 50
	fmt.Println("2   ->", int(2))   // 2
	fmt.Println(compress([]byte{'a', 'a', 'a', 'b', 'b', 'a', 'a'}))
}

func compress(chars []byte) int {
	n := len(chars)
	if n == 0 {
		return 0
	}

	write := 0 // Позиция для записи в массив
	count := 1 // Количество одинаковых символов подряд

	for i := 1; i <= n; i++ {
		// Если символ меняется или конец массива
		if i == n || chars[i] != chars[i-1] {
			chars[write] = chars[i-1] // записываем символ
			write++
			if count > 1 {
				digits := getDigitsBytes(count)
				// цифры уже в нужном порядке
				for _, d := range digits {
					chars[write] = d
					write++
				}
			}
			count = 1 // сбрасываем счётчик
		} else {
			count++
		}
	}

	return write
}

func getDigitsBytes(num int) []byte {
	if num == 0 {
		return []byte{'0'}
	}
	digits := make([]byte, 0)
	stack := make([]byte, 0)
	for num > 0 {
		stack = append(stack, '0'+byte(num%10))
		num /= 10
	}
	for i := len(stack) - 1; i >= 0; i-- {
		digits = append(digits, stack[i])
	}
	return digits
}
