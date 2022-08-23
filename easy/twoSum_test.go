package easy

import (
	"fmt"
	"testing"
)

// 枚举解题 时间复杂度O(n²) 空间复杂度O(1)
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

// hash表 空间复杂度O(n) 时间复杂度O(n)
func twnSumHash(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, x := range nums {
		if p, ok := hashTable[target-x]; ok {
			return []int{i, p}
		}
		hashTable[x] = i
	}
	return []int{}
}

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9

	res := twoSum(nums, target)
	res2 := twnSumHash(nums, target)
	fmt.Println(res, res2)
}
