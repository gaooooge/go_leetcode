package easy

import (
	"fmt"
	"testing"
)

// 寻找数组的中心下标 数组左右平衡， 感觉和双层遍历效率差不多少
func pivotIndex(n []int) int {
	total := 0
	for i := range n {
		total += n[i]
	}

	leftSum := 0
	rightSum := 0
	for i := range n {
		rightSum = total - leftSum - n[i]

		if leftSum == rightSum {
			return i
		}
		leftSum += n[i]
	}
	return -1
}

func TestPivotIndex(t *testing.T) {
	n := []int{-1, -1, -1, -1, -1, -1}

	res := pivotIndex(n)

	fmt.Println(res)
}
