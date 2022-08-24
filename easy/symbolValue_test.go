package easy

import (
	"fmt"
	"testing"
)

var symbol = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func TestSymbolValue(t *testing.T) {
	num := "MCMXCIV"
	res := symbolValue(num)
	fmt.Println("罗马文")
	fmt.Println(res)
}

func symbolValue(s string) int {
	len := len(s)
	num := 0
	for i := range s {
		value := symbol[s[i]]
		if i+1 < len && value < symbol[s[i+1]] {
			num -= value
		} else {
			num += value
		}

	}
	return num
}
