package randomPick

import "sort"

func backtracking(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	if target < 0 {
		return result
	}
	for i, item := range candidates {
		if i > 0 && item == candidates[i-1] {
			continue
		}
		tempA := []int{item}
		if target-item == 0 {
			result = append(result, tempA)
		} else {
			for _, tR := range backtracking(candidates[i+1:], target-item) {
				result = append(result, append(tempA, tR...))
			}
		}
	}
	return result
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	result := backtracking(candidates, target)
	return result
}
