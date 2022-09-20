package easy

import (
	"fmt"
	"math"
)

//设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈
type MinStack struct {
	nums    []int
	minNums []int
}

func Constructor() MinStack {
	n := MinStack{
		[]int{},
		[]int{math.MaxInt64},
	}
	return n
}

func min(n, m int) int {
	if n > m {
		return m
	}

	return n
}

func (this *MinStack) Push(val int) {
	this.nums = append(this.nums, val)
	top := this.minNums[len(this.minNums)-1]
	this.minNums = append(this.minNums, min(top, val))
}

func (this *MinStack) Pop() {
	this.nums = this.nums[:len(this.nums)-1]
	this.minNums = this.minNums[:len(this.minNums)-1]
}

func (this *MinStack) Top() int {
	return this.nums[len(this.nums)-1]
}

func (this *MinStack) GetMin() int {
	return this.minNums[len(this.minNums)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
