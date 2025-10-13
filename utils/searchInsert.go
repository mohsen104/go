package utils

import "slices"

func SearchInsert(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return i
		}
	}
	nums = append(nums, target)
	slices.Sort(nums)
	return slices.Index(nums, target)
}