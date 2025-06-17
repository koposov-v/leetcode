package main

func main() {
	nums := []int{1, 2, 87, 87, 87, 2, 1}
	candy(nums)
}
func candy(ratings []int) int {
	n := len(ratings)
	candies := make([]int, n)
	for i := range candies {
		candies[i] = 1
	}

	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}

	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			if candies[i] <= candies[i+1] {
				candies[i] = candies[i+1] + 1
			}
		}
	}

	totalCandies := 0
	for _, candy := range candies {
		totalCandies += candy
	}

	return totalCandies
}
