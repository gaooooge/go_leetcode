package easy

import (
	"fmt"
	"testing"
)

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
