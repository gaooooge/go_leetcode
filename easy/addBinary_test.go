package easy

import (
	"fmt"
	"strconv"
	"testing"
)

// 二进制加法，先从位数位加起
func addBinary(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}
	str := ""
	flag := 0
	lenAB := len(a) - len(b)

	for i := 0; i < lenAB; i++ {
		b = "0" + b
	}

	for i := len(a) - 1; i >= 0; i-- {
		aNum := a[i] - '0'
		bNum := b[i] - '0'

		sum := int(aNum+bNum) + flag
		if sum > 1 {
			str = strconv.Itoa(sum%2) + str
			flag = 1
		} else {
			str = strconv.Itoa(sum) + str
			flag = 0
		}
	}
	if flag > 0 {
		str = "1" + str
	}
	return str
}

func TestAddBinary(t *testing.T) {
	a := "1010"
	b := "1011"
	res := addBinary(a, b)
	fmt.Println(res)
}
