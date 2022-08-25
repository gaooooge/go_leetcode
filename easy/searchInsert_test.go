package easy

import (
	"fmt"
	"testing"
)

// // 遍历插入，如果不理想会时间复杂度O(n)
// func searchInsert(nums []int, target int) int {
// 	for i := 0; i < len(nums); i++ {
// 		if i < len(nums) && target <= nums[i] {
// 			return i
// 		}
// 	}
// 	return len(nums)
// }

// 二分法
func searchInsert(nums []int, target int) int {
	n := len(nums)
	left, right, ans := 0, n-1, n

	for left < right {
		fmt.Println(right, left, ((right - left) / 2))
		mid := (right-left)>>1 + left
		if target <= nums[mid] {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return ans
}

func TestSearchInsert(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	target := 5
	searchInsert(nums, target)
}
