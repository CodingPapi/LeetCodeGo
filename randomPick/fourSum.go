package randomPick

import (
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	result := make([][]int, 0)
	if len(nums) < 4 {
		return result
	}
	sort.Ints(nums)
	for index, first := range nums {
		if index > 0 && nums[index-1] == nums[index] {
			continue
		}
		for indexTail := len(nums) - 1; indexTail >= index+3; indexTail-- {
			if indexTail < len(nums)-1 && nums[indexTail] == nums[indexTail+1] {
				continue
			}
			neededSum := target - first - nums[indexTail]
			for j, k := index+1, indexTail-1; j < k; {
				if j > index+1 && nums[j] == nums[j-1] {
					j++
					continue
				}
				if k < indexTail-1 && nums[k] == nums[k+1] {
					k--
					continue
				}
				tempSum := nums[j] + nums[k]
				if tempSum == neededSum {
					result = append(result, []int{first, nums[j], nums[k], nums[indexTail]})
					j++
					k--
					continue
				} else if tempSum > neededSum {
					k--
					continue
				} else {
					j++
					continue
				}
			}

		}
	}
	return result
}
