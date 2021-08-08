package randomPick

func combinationSum(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	if target < 0 {
		return result
	}
	for i, item := range candidates {
		tempA := []int{item}
		if target-item == 0 {
			result = append(result, tempA)
		} else {
			for _, tR := range combinationSum(candidates[i:], target-item) {
				result = append(result, append(tempA, tR...))
			}
		}
	}
	return result
}
