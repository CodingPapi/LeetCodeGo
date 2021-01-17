package stackAndQueue

import (
	"math"
	// . "github.com/CodingPapi/LeetCodeGo/util"
)

// 155
type MinStack struct {
	diffStack []int
	min       int
}

func Constructor() MinStack {
	return MinStack{diffStack: make([]int, 0), min: math.MaxInt64}
}

func (this *MinStack) Push(x int) {
	if len(this.diffStack) == 0 {
		this.diffStack = append(this.diffStack, 0)
		this.min = x
		return
	}
	diff := x - this.min
	this.diffStack = append(this.diffStack, diff)
	if diff < 0 {
		this.min = x
	}
}

func (this *MinStack) Pop() {
	topDiff := this.diffStack[len(this.diffStack)-1]
	this.diffStack = this.diffStack[:len(this.diffStack)-1]
	if topDiff < 0 {
		this.min = this.min - topDiff
	}
}

func (this *MinStack) Top() int {
	topDiff := this.diffStack[len(this.diffStack)-1]
	if topDiff >= 0 {
		return topDiff + this.min
	} else {
		return this.min
	}
}

func (this *MinStack) GetMin() int {
	return this.min
}

// 155 end
