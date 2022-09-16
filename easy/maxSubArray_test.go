package easy

import (
	"fmt"
	"testing"
)

//最大子序和， 动态规划
func maxSubArray(n []int) int {
	max := n[0]
	for i := 1; i < len(n); i++ {
		if n[i]+n[i-1] > n[i] {
			n[i] += n[i-1] //存入当前值
		}

		if n[i] > max {
			max = n[i] //最大max
		}
	}
	return max
}

func TestMaxSubArray(t *testing.T) {
	n := []int{1, 2, 3, 0, -1}
	res := maxSubArray(n)
	fmt.Println(res)
}
