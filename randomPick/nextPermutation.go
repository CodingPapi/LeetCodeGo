package randomPick

func nextPermutation(nums []int) {
	endIndex := len(nums) - 1
	markIndex := endIndex - 1
	found := false
	for endIndex > 0 {
		markIndex = endIndex - 1
		for markIndex >= 0 {
			if nums[endIndex] > nums[markIndex] {
				nums[endIndex], nums[markIndex] = nums[markIndex], nums[endIndex]
				found = true
				break
			}
			markIndex--
		}
		if found {
			// markIndex + 1 -> endIndex
			for start, end := markIndex+1, len(nums)-1; start < end; {
				index := start
				for index < end {
					if nums[index] > nums[index+1] {
						nums[index], nums[index+1] = nums[index+1], nums[index]
					}
					index++
				}
				end--
			}
			break
		}
		endIndex--
	}
}
