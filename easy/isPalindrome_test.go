package easy

import (
	"fmt"
	"testing"
)

// 判断数字是否是回文
func isPalindrome(x int) bool {
	if x%10 == 0 && x != 0 {
		return false
	}

	rightNum := 0
	//避免超出最大长度 将原数据分割做对比
	for rightNum < x {
		rightNum = rightNum*10 + x%10 //右侧数据每次前移一位并拼接x个数位
		x /= 10                       //去掉原x个数位
	}

	return x == rightNum || x == rightNum/10
}

func TestIsPalindrome(t *testing.T) {
	x := 1234567654321
	fmt.Println("回文数")
	bool := isPalindrome(x)

	fmt.Println(bool)
}
