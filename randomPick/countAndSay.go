package randomPick

import "strconv"

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	tempStr := countAndSay(n - 1)
	result := ""
	tempChar := ' '
	currCount := 0
	for i, char := range tempStr {
		if char != tempChar {
			if i > 0 {
				result = result + strconv.Itoa(currCount) + string(tempChar)
			}
			tempChar = char
			currCount = 1
		} else {
			currCount++
		}
	}
	result = result + strconv.Itoa(currCount) + string(tempChar)
	return result
}
