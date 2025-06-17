package main

func majorityElement(nums []int) int {
	val, count := nums[0], 1

	for _, v := range nums {
		if count == 0 {
			val = v
			count = 1
		} else {
			if val == v {
				count++
			} else {
				count--
			}
		}
	}
	return val
}
