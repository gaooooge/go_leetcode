package easy

import (
	"fmt"
	"testing"
)

// 打家劫舍
func rob(nums []int) int {
	odd := 0
	even := 0

	for i := 0; i < len(nums); i++ {
		if i%2 == 1 { //even
			even += nums[i]
			even = max(odd, even)
		} else { //odd
			odd += nums[i]
			odd = max(odd, even)
		}
	}

	return max(odd, even)
}

func max(n, m int) int {
	if n > m {
		return n
	}
	return m
}

func TestRob(t *testing.T) {
	n := []int{1, 2, 3, 1}

	res := rob(n)

	fmt.Println(res)
}
