package stackAndQueue

import (
	"math"
	"strconv"
	// . "github.com/CodingPapi/LeetCodeGo/util"
)

//155
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

//155 end

//150
func EvalRPN(tokens []string) int {
	if len(tokens) < 3 {
		return 0
	}
	stack := make([]int, 0)
	for _, value := range tokens {
		switch value {
		case "+":
			{
				opLast := stack[len(stack)-1]
				opFormer := stack[len(stack)-2]
				stack = stack[:len(stack)-2]
				stack = append(stack, opFormer+opLast)
			}
		case "-":
			{
				opLast := stack[len(stack)-1]
				opFormer := stack[len(stack)-2]
				stack = stack[:len(stack)-2]
				stack = append(stack, opFormer-opLast)
			}
		case "*":
			{
				opLast := stack[len(stack)-1]
				opFormer := stack[len(stack)-2]
				stack = stack[:len(stack)-2]
				stack = append(stack, opFormer*opLast)
			}
		case "/":
			{
				opLast := stack[len(stack)-1]
				opFormer := stack[len(stack)-2]
				stack = stack[:len(stack)-2]
				stack = append(stack, opFormer/opLast)
			}
		default:
			{
				oper, _ := strconv.Atoi(value)
				stack = append(stack, oper)
			}
		}
	}

	return stack[len(stack)-1]
}

//394
type IntStack struct {
	value []int
}
type StringStack struct {
	value []string
}

func (this *IntStack) Push(x int) {
	this.value = append(this.value, x)
}

func (this *IntStack) Pop() int {
	top := this.value[len(this.value)-1]
	this.value = this.value[:len(this.value)-1]
	return top
}

func (this *IntStack) Top() int {
	top := this.value[len(this.value)-1]
	return top
}

func (this *StringStack) Push(x string) {
	this.value = append(this.value, x)
}

func (this *StringStack) Pop() string {
	top := this.value[len(this.value)-1]
	this.value = this.value[:len(this.value)-1]
	return top
}

func (this *StringStack) Top() string {
	top := this.value[len(this.value)-1]
	return top
}

func DecodeString(s string) string {
	kStack := &IntStack{value: []int{}}
	sStack := &StringStack{value: []string{""}}
	kValue := ""
	for _, value := range s {
		switch {
		case value == '[':
			// "["
			{
				repeatK, _ := strconv.Atoi(kValue)
				kStack.Push(repeatK)
				kValue = ""
				sStack.Push("")
			}
		case value == ']':
			// "]"
			{
				latest := sStack.Pop()
				repeatK := kStack.Pop()

				result := ""
				for i := 0; i < repeatK; i++ {
					result += latest
				}
				latest = sStack.Pop()
				result = latest + result
				sStack.Push(result)
			}
		case value >= '0' && value <= '9':
			// number
			{
				kValue += string(value)
			}
		default:
			// char
			{
				latest := sStack.Pop()
				latest += string(value)
				sStack.Push(latest)
			}
		}

	}

	result := ""
	for len(sStack.value) > 0 {
		top := sStack.Pop()
		if len(top) > 0 {
			result = top + result
		}
	}
	return result
}
