package easy

import "testing"

func plusOne(d []int) []int {
	j := 1
	for i := len(d) - 1; i >= 0; i-- {
		if d[i]+j >= 10 {
			j = 1
			d[i] = (d[i] + j) % 10
		} else {
			d[i] = d[i] + j
			j = 0
		}
	}
	if j != 0 {
		d = append([]int{j}, d...)
	}

	return d
}

func TestPlusOne(t *testing.T) {
	digits := []int{1, 2, 3, 4, 5}
	plusOne(digits)
}
