package easy

import (
	"fmt"
	"testing"
)

// // 不知道错在哪，看解题需要使用双指针
// func removeDuplicates(nums []int) int {
// 	hashTable := map[int]bool{}
// 	for i, v := range nums {
// 		if ok := hashTable[v]; !ok {
// 			hashTable[v] = true
// 		} else {
// 			nums = append(nums[i:], nums[i+1:]...)
// 		}
// 	}
// 	return len(hashTable)
// }

func removeDuplicatesV2(nums []int) int {
	index1, index2 := 0, 0
	for index2 < len(nums) {
		if nums[index2] <= nums[index1] {
			index2++
		} else {
			index1++
			nums[index1] = nums[index2]
			index2++
		}
	}
	fmt.Println(nums)
	return index1 + 1
}

func TestRemoveDuplicates(t *testing.T) {
	arr := []int{1, 1, 2, 3, 4, 5, 6, 7, 7, 7}
	count := removeDuplicatesV2(arr)
	fmt.Println("删除有序数组中的重复项")
	fmt.Println(count)
}
