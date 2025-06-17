package main

import "fmt"

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := (left + right) / 2
		cur := nums[mid]
		leftValue := nums[left]
		rightValue := nums[right]

		if cur == target {
			return mid
		}

		if leftValue <= cur {
			if leftValue <= target && target < cur {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if cur < target && cur <= rightValue {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}

type Res struct {
	input  []int
	target int
	res    int
}

func main() {
	testCases := []Res{
		{
			[]int{4, 5, 6, 7, 0, 1, 2},
			0,
			4,
		},
		{
			[]int{4, 5, 6, 7, 0, 1, 2},
			3,
			-1,
		},
		{
			[]int{1},
			0,
			-1,
		},
	}
	for i, tt := range testCases {
		res := search(tt.input, tt.target)
		if res != tt.res {
			fmt.Printf("Test-%d, Ожидали %d, а получили %d\n", i, tt.res, res)
		}
	}
}
