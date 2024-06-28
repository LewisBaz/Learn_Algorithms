package twopointers

// https://leetcode.com/problems/remove-element/?envType=study-plan-v2&envId=top-interview-150

func removeElement(nums []int, val int) int {
	startPointer := 0
	endPointer := len(nums) - 1
	for startPointer <= endPointer {
		if nums[endPointer] == val {
			endPointer -= 1
		} else {
			nums[startPointer], nums[endPointer] = nums[endPointer], nums[startPointer]
			startPointer += 1
		}
	}
	nums = nums[:endPointer+1]
	return len(nums)
}