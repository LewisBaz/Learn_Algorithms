package arrays

import (
	"fmt"
)

func TwoSum(nums []int, target int) []int {
    for i := 0; i < len(nums) - 1; i++ {
		for j := range nums {
			if i != j && nums[i] + nums[j] == target {
				fmt.Println(i, j)
				return []int{i, j}
			}
		}
	}
	return []int{}
}