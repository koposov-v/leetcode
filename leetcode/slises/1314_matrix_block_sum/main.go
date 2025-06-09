package main

import "fmt"

func main() {
	mat := [][]int{{67, 64, 78}, {99, 98, 38}, {82, 46, 46}, {6, 52, 55}, {55, 99, 45}}
	fmt.Println(matrixBlockSum(mat, 3))
}

func matrixBlockSum(mat [][]int, k int) [][]int {
	rows := len(mat) + 1
	cols := len(mat[0]) + 1

	px := make([][]int, rows)
	for i := range rows {
		px[i] = make([]int, cols)
	}

	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			pxSum := px[i-1][j] + px[i][j-1] - px[i-1][j-1]
			px[i][j] = pxSum + mat[i-1][j-1]
		}
	}

	for i := 0; i < rows-1; i++ {
		for j := 0; j < cols-1; j++ {
			r1 := max(i-k, 0)
			r2 := min(i+k+1, rows-1)
			c1 := max(j-k, 0)
			c2 := min(j+k+1, cols-1)

			mat[i][j] = px[r2][c2] - px[r1][c2] - px[r2][c1] + px[r1][c1]
		}
	}

	return mat
}

func missingNumber(nums []int) int {
	n := len(nums)
	arrSum := n * (n + 1) / 2

	for _, v := range nums {
		arrSum -= v
	}

	return arrSum
}
