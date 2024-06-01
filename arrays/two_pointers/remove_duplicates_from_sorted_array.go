package twopointers

// https://leetcode.com/problems/remove-duplicates-from-sorted-array/description/

func RemoveDuplicates(nums []int) int {
	curr := 0
	next := curr + 1
	for next < len(nums) {
		if nums[curr] == nums[next] {
			nums = append(nums[:curr], nums[next:]...)
		} else {
			curr += 1
			next += 1
		}
	}
	return len(nums)
}