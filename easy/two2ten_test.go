package easy

import (
	"fmt"
	"math"
	"testing"
)

func TestTwo2Ten(t *testing.T) {
	two2Ten()
}

// 突然被问到了二级制到十进制的转换，学习一下
func two2Ten() {
	x := 101101
	i := 0
	num := 0
	for x > 0 {
		n := float64(x % 10)
		num += int(math.Pow(n*2, float64(i)))
		x /= 10
		i++
	}
	fmt.Println("进制转换")
	fmt.Println(num)
}
