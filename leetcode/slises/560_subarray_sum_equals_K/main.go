package main

/**
Суть идеи в том, если текущая префиксная сумма минус k уже была, значит, между этими точками есть подмассив с суммой k.
*/

func subarraySum(nums []int, k int) int {
	prefixSum := make(map[int]int)
	prefixSum[0] = 1

	currSum := 0
	count := 0

	for _, v := range nums {
		currSum += v
		if _, ok := prefixSum[currSum-k]; ok {
			count += prefixSum[currSum-k]
		}

		prefixSum[currSum] += 1
	}

	return count
}
