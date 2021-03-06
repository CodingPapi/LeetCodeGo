package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello Wold!")
	nums := [][]int{}
	fmt.Println(len(nums))
}

func backtrace(nums []int, currPos int, path []int, result *[][]int) {
	if currPos >= len(nums) {
		ans := make([]int, len(path))
		copy(ans, path)
		*result = append(*result, ans)
		return
	}
	path = append(path, nums[currPos])
	backtrace(nums, currPos+1, path, result)
	path = path[0 : len(path)-1]
	backtrace(nums, currPos+1, path, result)
}

func backtrace2(nums []int, currPos int, path []int, result *[][]int) {
	ans := make([]int, len(path))
	copy(ans, path)
	*result = append(*result, ans)
	for i := currPos; i < len(nums); i++ {
		path = append(path, nums[i])
		backtrace(nums, i+1, path, result)
		path = path[0 : len(path)-1]
	}
}
