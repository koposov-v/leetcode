package main

func main() {
	subarraySum([]int{1, 1, 1}, 2)
}
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
