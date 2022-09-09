package easy

import (
	"fmt"
	"strings"
	"testing"
)

func isPalindromeStr(s string) bool {
	//先全部转换成小写
	s = strings.ToLower(s)
	//定义两个指针分别从字符串两头对比
	left, right := 0, len(s)-1
	for left < right {
		if !isalnum(s[left]) { //左侧不在区间内跳过进一
			left++
			continue
		}
		if !isalnum(s[right]) { //右侧不在区间内跳过退一
			right--
			continue
		}
		if s[left] != s[right] { //如果左右不相同则直接返回false
			return false
		}
		left++
		right--
	}
	return true
}

// 方法用于判断是否在对比区间
func isalnum(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}

func TestIsPalindromeStr(t *testing.T) {
	s := "A man, a plan, a canal: Panama"
	res := isPalindromeStr(s)

	fmt.Println(res)
}
