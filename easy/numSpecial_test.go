package easy

import "testing"

// 二进制矩阵中的特殊位置
func numSpecial(mat [][]int) (ans int) {
	rowsSum := make([]int, len(mat))
	colsSum := make([]int, len(mat[0]))
	for i, row := range mat {
		for j, x := range row {
			rowsSum[i] += x
			colsSum[j] += x
		}
	}
	for i, row := range mat {
		for j, x := range row {
			if x == 1 && rowsSum[i] == 1 && colsSum[j] == 1 {
				ans++
			}
		}
	}
	return
}

func TestNumSpecial(t *testing.T) {
	mat := [][]int{
		{0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0},
		{0, 1, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 1, 1},
	}

	numSpecial(mat)
}
