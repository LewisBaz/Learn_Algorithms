package binarysearch

// import "fmt"

// https://leetcode.com/problems/search-insert-position/description/

func searchInsert(nums []int, target int) int {
    left := 0
	right := len(nums) - 1
	var possible_index int
	for left <= right {
		mid := left + (right - left) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
			possible_index = mid + 1
		} else {
			right = mid - 1
			possible_index = mid
		}
	}
	return possible_index 
}