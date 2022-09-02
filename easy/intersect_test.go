package easy

import (
	"fmt"
	"testing"
)

func intersect(nums1, nums2 []int) []int {
	cMap := map[int]int{}
	res := []int{}
	for i := range nums1 {
		cMap[nums1[i]] += 1
	}

	for i := range nums2 {
		if c, ok := cMap[nums2[i]]; ok && c > 0 {
			cMap[nums2[i]] -= 1
			fmt.Println(cMap)
			res = append(res, nums2[i])
		}
	}

	return res
}

func TestIntersect(t *testing.T) {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}

	res := intersect(nums1, nums2)
	fmt.Println(res)
}
