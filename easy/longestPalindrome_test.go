package easy

import (
	"fmt"
	"testing"
)

// 最长回文数
func longestPalindrome(s string) int {
	m := make(map[int32]int, 52)
	for _, j := range s {
		m[j]++
	}
	sum := 0
	flag := 0
	for _, v := range m {
		if v%2 == 0 {
			sum += v
		} else {
			sum += v - 1
			flag = 1
		}
	}
	return sum + flag
}

func TestLongestPalindrome(t *testing.T) {
	s := "abccccdd"
	res := longestPalindrome(s)
	fmt.Println(res)
}
