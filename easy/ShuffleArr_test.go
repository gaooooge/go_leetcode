package easy

import "math/rand"

//设计问题、打乱数组
type Solution struct {
	original []int
	nums []int
}


func Constructor(s []int) Solution {
	return Solution{original: s, nums: append([]int(nil), len(s))}
}

func (s *Solution) Reset() []int {
	copy(s.nums, s.original)

	return s.nums
}

//Fisher-Yates洗牌算法
func (s *Solution) Shuffle() []int {
	n := len(s.nums)
	for i := range s.nums {
		j := i + rand.Intn(n - 1)
		s.nums[i], s.nums[j] = s.nums[j], s.nums[i]
	}

	return s.nums
}