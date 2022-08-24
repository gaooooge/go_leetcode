package easy

import (
	"fmt"
	"testing"
)

var kh = map[byte]byte{
	')': '(',
	'}': '{',
	']': '[',
}

// 有效的括号
func isValid(s string) bool {
	//如果字符串的长度是基数则必然不可能形成有效括号
	if len(s)%2 > 0 {
		return false
	}
	fmt.Println(kh['<'])
	stack := []byte{}
	for i := range s {
		if kh[s[i]] > 0 { //如果取到了右括号
			if len(stack) == 0 || stack[len(stack)-1] != kh[s[i]] { //数值已经为空或者最后一位不是与之相对应的括号
				return false
			}
			stack = stack[:len(stack)-1]
		} else { //取不到右括号时将当前的值写入stack[
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}

func TestIsValid(t *testing.T) {
	str := "[[]][[]](())"
	res := isValid(str)
	fmt.Println("有效的括号")
	fmt.Println(res)
}
