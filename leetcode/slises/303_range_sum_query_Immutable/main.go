package main

type NumArray struct {
	nums       []int
	prefixNums []int
}

func Constructor(nums []int) NumArray {
	prefixNums := make([]int, 1, len(nums)+1)

	for _, v := range nums {
		prefVal := prefixNums[len(prefixNums)-1] + v
		prefixNums = append(prefixNums, prefVal)
	}

	return NumArray{
		nums:       nums,
		prefixNums: prefixNums,
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.prefixNums[right+1] - this.prefixNums[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
