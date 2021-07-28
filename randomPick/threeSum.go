package randomPick

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	for index, first := range nums {
		if first > 0 {
			break
		}
		if index > 0 && nums[index-1] == nums[index] {
			continue
		}
		neededSum := -1 * first
		for j := 1; j < len(nums)-index-1; j++ {
			if j > 1 && nums[index+j-1] == nums[index+j] {
				continue
			}
			k := len(nums) - j
			tempSum := 0
			for k > index+j {
				if j > 1 && nums[index+j-1] == nums[index+j] {
					j++
					continue
				}
				tempSum = nums[index+j] + nums[k]
				if tempSum > neededSum {
					k--
					continue
				} else if tempSum < neededSum {
					j++
					continue
				}
				result = append(result, []int{first, nums[index+j], nums[k]})
				j++
			}
			break
		}
	}
	return result
}
