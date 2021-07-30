package randomPick

func generateParenthesis(n int) []string {
	result, _ := generate("", n, n)
	return result
}

func generate(currentString string, xLeft int, yLeft int) (currResult []string, resultOk bool) {
	if yLeft < xLeft {
		return nil, false
	}
	if xLeft == 0 && yLeft == 0 {
		return []string{currentString}, true
	}
	result := []string{}
	if xLeft > 0 {
		leftResult, isOk := generate(currentString+"(", xLeft-1, yLeft)
		if isOk {
			result = append(result, leftResult...)
		}
	}
	if yLeft > 0 {
		rightResult, isOk := generate(currentString+")", xLeft, yLeft-1)
		if isOk {
			result = append(result, rightResult...)
		}
	}
	return result, true
}
