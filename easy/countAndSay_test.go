package easy

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func countAndSay(n int) string {
	prev := "1"
	for i := 2; i <= n; i++ {
		cur := &strings.Builder{}
		for j, start := 0, 0; j < len(prev); start = j {
			for j < len(prev) && prev[j] == prev[start] {
				j++
			}
			cur.WriteString(strconv.Itoa(j - start))
			cur.WriteByte(prev[start])
		}
		prev = cur.String()
	}
	return prev
}

func TestCountAndSay(t *testing.T) {
	res := countAndSay(4)

	fmt.Println(res)
}
