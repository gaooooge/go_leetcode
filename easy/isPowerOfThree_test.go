package easy

import (
	"fmt"
	"testing"
)

// 3的幂
func isPowerOfThree(n int) bool {
	return 1162261467%n == 0
}

func TestIsPowerOfThree(t testing.T) {
	n := 9
	res := isPowerOfThree(n)
	fmt.Println(res)
}
