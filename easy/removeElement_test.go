package easy

import (
	"fmt"
	"testing"
)

func removeElement(nums []int, val int) int {
	num := 0
	for i, v := range nums {
		if nums[i] != val {
			nums[num] = v
			num++
		}
	}
	fmt.Println(nums)
	return num
}

func TestRemoveElement(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	val := 2
	res := removeElement(nums, val)
	fmt.Println("移除元素")
	fmt.Println(res)
}
