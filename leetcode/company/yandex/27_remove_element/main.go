package main

import "fmt"

func main() {
	test := []int{1}
	fmt.Println(removeElement(test, 1))
}
func removeElement(nums []int, val int) int {
	i := 0
	for _, v := range nums {
		if v != val {
			nums[i] = v
			i++
		}
	}
	return i
}
