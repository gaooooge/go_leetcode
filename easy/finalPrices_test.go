package easy

import (
	"fmt"
	"testing"
)

func finalPrices(p []int) []int {
	for i := 0; i < len(p); i++ {
	OutFor:
		for j := i + 1; j < len(p); j++ {
			if p[i] >= p[j] {
				p[i] = p[i] - p[j]
				break OutFor
			}
		}

	}

	return p
}

func TestFinalPrices(t *testing.T) {
	p := []int{8, 4, 6, 2, 3}
	res := finalPrices(p)
	fmt.Println(res)
}
