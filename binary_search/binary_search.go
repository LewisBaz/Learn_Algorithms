package binarysearch

func Search(nums []int, target int) int {
	origin := nums
	counter := 0
	var index int = -1
	for index == -1 {
		length := len(origin)
		left := origin[:length/2]
		right := origin[length/2:]
		if len(left) != 0 && target > left[len(left)-1] {
			origin = right
			counter += len(left)
		} else if len(left) != 0 && target < left[len(left)-1] {
			origin = left
		} else if len(left) != 0 {
			counter += len(left)
			index = counter - 1
		} else if counter == len(nums) - 1 && nums[counter] == target {
			index = counter
		} else {
			return index
		}
	}
	return index
}