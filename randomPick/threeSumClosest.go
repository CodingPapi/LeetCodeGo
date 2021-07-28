package randomPick

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	result := math.MaxInt32
	tempSum := 0
	for i := 0; i < len(nums)-2; i++ {
		j, k := i+1, len(nums)-1
		for k > j {
			tempSum = nums[i] + nums[j] + nums[k]
			if absInt(tempSum-target) < absInt(result-target) {
				result = tempSum
			}
			if tempSum > target {
				k--
			} else if tempSum < target {
				j++
			} else {
				return target
			}
		}
	}
	return result
}

func absInt(t int) int {
	if t < 0 {
		return -t
	}
	return t
}
