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
	for index, indexTail := 0, len(nums)-1; indexTail >= index+3; {
		if index > 0 && nums[index-1] == nums[index] {
			index++
			continue
		}
		if indexTail < len(nums)-1 && nums[indexTail+1] == nums[indexTail] {
			indexTail--
			continue
		}
		neededSum := target - nums[index] - nums[indexTail]
		lastSum := 0
		for j, k := index+1, indexTail-1; j < k; {
			if j > index+1 && nums[j] == nums[j-1] {
				j++
				continue
			}
			if k < indexTail-1 && nums[k] == nums[k+1] {
				k--
				continue
			}
			lastSum = nums[j] + nums[k]
			if lastSum == neededSum {
				result = append(result, []int{nums[index], nums[j], nums[k], nums[indexTail]})
				j++
				k--
				continue
			} else if lastSum > neededSum {
				k--
				continue
			} else {
				j++
				continue
			}
		}
		if lastSum > neededSum {
			indexTail--
		} else {
			index++
		}
	}
	return result
}
