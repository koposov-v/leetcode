package main

/**
Здесь мы
*/

type NumMatrix struct {
	px [][]int
}

// time: O(n * m)
// memory: O(n * m)

func Constructor(matrix [][]int) NumMatrix {
	rows := len(matrix)
	cols := len(matrix[0])
	px := make([][]int, rows+1)
	for i := range px {
		px[i] = make([]int, cols+1)
	}

	for i := 1; i < rows+1; i++ {
		for j := 1; j < cols+1; j++ {
			idxSum := px[i][j-1] + px[i-1][j] - px[i-1][j-1]
			px[i][j] = matrix[i-1][j-1] + idxSum
		}
	}

	return NumMatrix{
		px: px,
	}
}

//time: O(1)
//memory: O(1)

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	col2++
	row2++

	return this.px[row2][col2] - this.px[row1][col2] - this.px[row2][col1] + this.px[row1][col1]
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
