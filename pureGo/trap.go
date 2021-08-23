package pureGo

type Stack struct {
	holder []int
}

func (s *Stack) Init()    { s.holder = make([]int, 0) }
func (s *Stack) Len() int { return len(s.holder) }
func (s *Stack) Push(item int) {
	s.holder = append(s.holder, item)
}
func (s *Stack) Pop() int {
	if len(s.holder) > 0 {
		result := s.holder[len(s.holder)-1]
		s.holder = s.holder[:len(s.holder)-1]
		return result
	} else {
		return -1
	}
}
func (s *Stack) Peek() int {
	if len(s.holder) > 0 {
		return s.holder[len(s.holder)-1]
	} else {
		return -1
	}
}

func min(a int, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func trap(height []int) int {
	var stack Stack
	stack.Init()
	sum := 0
	for i, item := range height {
		for stack.Len() > 0 && item > height[stack.Peek()] {
			floreIndex := stack.Pop()
			if stack.Len() <= 0 {
				break
			}
			tempHeight := min(item, height[stack.Peek()])
			sum += (tempHeight - height[floreIndex]) * (i - floreIndex)
		}
		stack.Push(i)
	}
	return sum
}
