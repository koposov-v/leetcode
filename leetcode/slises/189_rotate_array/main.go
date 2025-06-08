package main

func rotate(nums []int, k int) {
	k %= len(nums)
	nums = rotateArr(nums)
	nums = append(rotateArr(nums[:k]), rotateArr(nums[k:])...)
}

func rotateArr(nums []int) []int {
	i := 0
	j := len(nums) - 1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
	return nums
}
