package easy

import (
	"fmt"
	"testing"
)

func hammingWeight(num uint32) int {
	for i := 0; i < 32; i++ {
		if num > 0 && num<<1 == 1 {
			fmt.Println(1)
		}
	}
	return 1
}

func TestHammingWeight(t *testing.T) {
	var num uint32 = 11111111111111111111111111111101
	res := hammingWeight(num)

	fmt.Println(res)
}
