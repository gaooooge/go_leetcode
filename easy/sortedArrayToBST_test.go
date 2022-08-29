package easy

import (
	"fmt"
	"testing"
)

type treeNode3 struct {
	Val   int //
	Left  *treeNode3
	Right *treeNode3
}

func sortedArrayToBST(nums []int) *treeNode3 {
	length := len(nums)
	mix := length >> 1

	return &treeNode3{
		nums[mix],
		sortedArrayToBST(nums[:length]),
		sortedArrayToBST(nums[length+1:]),
	}
}

func TestSortedArrayToBST(t *testing.T) {
	nums := []int{-10, -3, 0, 5, 9}
	res := sortedArrayToBST(nums)

	fmt.Println(res)
}
