package main

import (
	"fmt"
)

func main() {

	fmt.Println(findLengthOfLCIS([]int{2, 1, 3}))
}

func findLengthOfLCIS(nums []int) int {
	ans := 1
	cur := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			cur++
			ans = max(ans, cur)
		} else {
			cur = 1
		}
	}

	return ans
}
