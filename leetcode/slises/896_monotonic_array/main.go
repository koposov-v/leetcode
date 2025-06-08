package main

func isMonotonic(nums []int) bool {
	isDec, isInc := true, true

	for i := 1; i < len(nums); i++ {

		prev := nums[i-1]
		cur := nums[i]
		isDec = isDec && prev >= cur
		isInc = isInc && prev <= cur
	}

	return isDec || isInc
}
