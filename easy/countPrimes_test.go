package easy

import (
	"fmt"
	"testing"
)

func isPrime(x int) bool {
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func countPrimes(n int) (cnt int) {
	for i := 2; i < n; i++ {
		if isPrime(i) {
			cnt++
		}
	}
	return
}

func TestCountPrimes(t *testing.T) {
	res := countPrimes(5)

	fmt.Println(res)
}
