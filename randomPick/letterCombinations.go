package randomPick

import "strconv"

var numsString [][]string = [][]string{
	{"0", "0", "0"},
	{"0", "0", "0"},
	{"a", "b", "c"},
	{"d", "e", "f"},
	{"g", "h", "i"},
	{"j", "k", "l"},
	{"m", "n", "o"},
	{"p", "q", "r", "s"},
	{"t", "u", "v"},
	{"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	if len(digits) < 1 {
		return []string{}
	}
	if len(digits) == 1 {
		index, _ := strconv.Atoi(string(digits[0]))
		return numsString[index]
	}
	currDigit, _ := strconv.Atoi(string(digits[0]))
	currDigits := numsString[currDigit]
	gotCombinations := letterCombinations(digits[1:])
	result := []string{}
	for _, digit := range currDigits {
		for _, comb := range gotCombinations {
			result = append(result, digit+comb)
		}
	}
	return result
}
