package sorting

import "fmt"

func SortColors(nums []int)  {
	len := len(nums)
	for i := range nums {
		if i + 1 < len {
			left := nums[i]
			right := nums[i+1]
			index := i
			for left >= right {
				nums[index] = right
				nums[index+1] = left
				index -= 1
				if index == -1 {
					break
				} else {
					left = nums[index]
					right = nums[index+1]
				}
			}
		}
	}
	fmt.Println(nums)
}