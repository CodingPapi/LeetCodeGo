package randomPick

func nextPermutation(nums []int) {
	endIndex := len(nums) - 1
	markIndex := endIndex - 1
	currMarkIndex := -1
	currTargetIndex := len(nums) - 1
	for endIndex > 0 {
		markIndex = endIndex - 1
		for markIndex >= 0 {
			if nums[endIndex] > nums[markIndex] && markIndex > currMarkIndex {
				currMarkIndex = markIndex
				currTargetIndex = endIndex
				break
			}
			markIndex--
		}
		endIndex--
	}
	if currMarkIndex > -1 {
		nums[currTargetIndex], nums[currMarkIndex] = nums[currMarkIndex], nums[currTargetIndex]
		bubbleSort(nums, currMarkIndex+1, len(nums)-1)
	} else {
		bubbleSort(nums, 0, len(nums)-1)
	}
}

func bubbleSort(nums []int, start int, end int) {
	for start < end {
		index := start
		for index < end {
			if nums[index] > nums[index+1] {
				nums[index], nums[index+1] = nums[index+1], nums[index]
			}
			index++
		}
		end--
	}
}
