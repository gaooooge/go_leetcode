package easy

import (
	"fmt"
	"testing"
)

// //排序遍历对比
// func singleNumber(nums []int) int {
// 	sort.Ints(nums)
// 	for i := 1; i < len(nums); i++ {
// 		if i == len(nums)-1 && nums[i] != nums[i-1] {
// 			return nums[i]
// 		} else if nums[i] != nums[i-1] && nums[i] != nums[i+1] {
// 			fmt.Println(nums, i)
// 			return nums[i]
// 		}
// 	}
// 	return nums[0]
// }

// 位运算对比
func singleNumber(nums []int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		result ^= nums[i]
		fmt.Println(result)
	}

	return result
}

func TestSingleNumber(t *testing.T) {
	res := singleNumber([]int{4, 1, 2, 1, 2})

	fmt.Println(res)
}
