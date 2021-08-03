package randomPick

func divide(dividend int, divisor int) int {
	maxInt := 1<<31 - 1
	minInt := -maxInt - 1
	symbol := 1
	if divisor == 0 || dividend == 0 {
		return 0
	}
	if divisor == 1 {
		return dividend
	}
	if divisor == -1 {
		if dividend == minInt {
			return maxInt
		} else {
			return -dividend
		}
	}
	if dividend < 0 {
		symbol = -symbol
		dividend = -dividend
	}
	if divisor < 0 {
		symbol = -symbol
		divisor = -divisor
	}

	result := binaryDivide(dividend, divisor, 0)
	if symbol == -1 {
		return 0 - result
	}
	return result
}

func binaryDivide(dividend int, divisor int, currQuo int) int {
	if dividend < divisor {
		return currQuo
	}
	quo := 1
	tempDivisor := divisor << 1
	for dividend > tempDivisor {
		quo = quo << 1
		tempDivisor = tempDivisor << 1
	}
	leftQuo := binaryDivide(dividend-tempDivisor>>1, divisor, currQuo+quo)
	return leftQuo
}
