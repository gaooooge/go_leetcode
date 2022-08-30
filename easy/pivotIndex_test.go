package easy

import (
	"fmt"
	"testing"
)

func pivotIndex(n []int) int {
	total := 0
	for i := range n {
		total += n[i]
	}

	leftSum := 0
	rightSum := 0
	for i := range n {
		rightSum = total - leftSum - n[i]
		fmt.Println(i, total, leftSum, rightSum)

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
