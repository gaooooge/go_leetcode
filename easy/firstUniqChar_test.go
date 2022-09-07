package easy

import (
	"fmt"
	"testing"
)

// 字符串中的第一个唯一字符
func firstUniqChar(s string) int {
	cnt := [26]int{}
	for _, ch := range s {
		cnt[ch-'a']++
	}
	for i, ch := range s {
		if cnt[ch-'a'] == 1 {
			return i
		}
	}
	return -1
}

func TestFirstUniqChar(t *testing.T) {
	s := "leetcode"
	res := firstUniqChar(s)

	fmt.Println(res)
}
