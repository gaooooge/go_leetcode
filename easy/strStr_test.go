package easy

import (
	"fmt"
	"testing"
)

// 实现strstr
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	for i := range haystack {
		if (i+len(needle) < len(haystack)) && haystack[i:i+len(needle)] == needle {
			return i
		}
	}

	return -1
}

func TestStrStr(t *testing.T) {
	haystack := "hello"
	needle := "ll"
	res := strStr(haystack, needle)
	fmt.Println("实现strstr")
	fmt.Println(res)
}
