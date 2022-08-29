package easy

import (
	"fmt"
	"sort"
	"testing"
)

func merge(num1 []int, m int, num2 []int, n int) []int {
	copy(num1[m:], num2)
	sort.Ints(num1)
	return num1
}

func TestMerge(t *testing.T) {
	num1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	num2 := []int{2, 5, 6}
	n := 3

	res := merge(num1, m, num2, n)
	fmt.Println(res)
}
