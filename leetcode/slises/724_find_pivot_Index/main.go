package main

/*
Здесь используется паттерн AllSum, так же можно решить задачу с помощью
префиксный/постфиксных массивов
*/

func pivotIndex(nums []int) int {
	allSum := 0
	for _, v := range nums {
		allSum += v
	}
	pxSum := 0
	for i, v := range nums {
		if allSum-v-pxSum == pxSum {
			return i
		}
		pxSum += v
	}

	return -1
}
