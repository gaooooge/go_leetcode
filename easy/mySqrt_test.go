package easy

import (
	"fmt"
	"testing"
)

func mySqrt(x int) int {
	if x == 1 {
		return 1
	}
	for i := 0; i < x; i++ {
		if i*i <= x && (i+1)*(i+1) > x {
			return i
		}
	}

	return 0
}

func TestMySqrt(t *testing.T) {
	x := 8
	res := mySqrt(x)
	fmt.Println(res)
}
