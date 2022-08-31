package medium

import (
	"fmt"
	"testing"
)

func validateStackSequences(pushed []int, popped []int) bool {
	if len(pushed) == 0 && len(popped) == 0 {
		return true
	}

	nums := []int{}
	for len(pushed) > 0 || len(popped) > 0 {

		if len(nums) == 0 && len(pushed) > 0 {
			nums = append(nums, pushed[0])
			pushed = pushed[1:]
		} else if len(popped) > 0 && nums[len(nums)-1] == popped[0] {
			nums = nums[:len(nums)-1]
			popped = popped[1:]
		} else if len(pushed) > 0 {
			nums = append(nums, pushed[0])
			pushed = pushed[1:]
		} else {
			return false
		}
		fmt.Println(pushed, nums, popped)
	}

	if len(pushed) == 0 && len(popped) == 0 {
		return true
	}

	return false
}

func TestValidateStackSequences(t *testing.T) {
	pushed := []int{1, 2, 3, 4, 5}
	popped := []int{4, 3, 5, 1, 2}
	res := validateStackSequences(pushed, popped)
	fmt.Println(res)
}
