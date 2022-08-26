package easy

import (
	"fmt"
	"testing"
)

// 原本还有额外的一个参数来统计是否重置，看了扣友的解法，自叹不如
func lengthOfLastWord(s string) int {
	num := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			num++
		} else if num != 0 {
			return num
		}
	}
	return num
}

func TestLengthOfLastWord(t *testing.T) {
	s := "  fly me   to   the moon  "
	res := lengthOfLastWord(s)
	fmt.Println(res)
}
