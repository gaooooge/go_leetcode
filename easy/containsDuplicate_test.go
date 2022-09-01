package easy

import (
	"fmt"
	"testing"
)

func containsDuplicate(n []int) bool {
	pushList := map[int]bool{}

	for i := range n {
		if ok := pushList[n[i]]; ok {
			return true
		} else {
			pushList[n[i]] = true
		}
	}

	return false
}

func TestContainsDuplicate(t *testing.T) {
	nums := []int{1, 2, 3, 4, 3}

	res := containsDuplicate(nums)

	fmt.Println(res)
}
