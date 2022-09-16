package easy

import (
	"fmt"
	"testing"
)

// 计算股票的最大收益
func maxProfit(p []int) int {
	minV := p[0]
    maxV := 0
	for i := 1; i < len(p); i++ {
        if p[i] < minV {
            minV = p[i]
        }

        if p[i] - minV > maxV {
            maxV = p[i]-minV
        }
	}
	return maxV
}

func TestMaxProfit(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 8}

	res := maxProfit(arr)

	fmt.Println(res)
}
