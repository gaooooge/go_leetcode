package easy

import (
	"fmt"
	"testing"
)

// 爬楼梯 滚动实现
func climbStairs(n int) int {
	q, p, r := 0, 0, 1

	for i := 0; i < n; i++ {
		p = q
		q = r
		r = q + p
	}

	return r
}

func TestClimbStairs(t *testing.T) {
	n := 10
	res := climbStairs(n)
	fmt.Println(res)
}
