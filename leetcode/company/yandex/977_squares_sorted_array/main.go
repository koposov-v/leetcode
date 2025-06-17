package main

func main() {
	sortedSquares([]int{-4, -1, 0, 3, 10})
}

func sortedSquares(nums []int) []int {
	n := len(nums)
	l, r := 0, n-1
	ptr := n - 1
	res := make([]int, n)
	for l <= r {
		lv := nums[l] * nums[l]
		rv := nums[r] * nums[r]
		if lv < rv {
			res[ptr] = rv
			r--
			ptr--
		} else {
			res[ptr] = lv
			l++
			ptr--
		}
	}

	return res
}
