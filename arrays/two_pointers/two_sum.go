package twopointers

// https://leetcode.com/problems/two-sum/description/

import (
	"fmt"
	"sort"
	"strconv"
)

func TwoSumBruteForce(nums []int, target int) []int {
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

func TwoSumHashMap(nums []int, target int) []int {
    dict := make(map[int]string)
	for i, n := range nums {
		if dict[target-n] != "" {
			ind, err := strconv.Atoi(dict[target-n])
			if err == nil  {
				return []int{i, ind}
			}
		}
		dict[n] = strconv.Itoa(i)
	} 
	return []int{}
}

// Подходит только для отсортированного массива
func TwoSumTwoPointers(nums []int, target int) []int {
	sort.Ints(nums)
	begin := 0
	end := len(nums) - 1
	for begin < end {
		if nums[begin] + nums[end] > target {
			end -= 1
		} else if nums[begin] + nums[end] < target {
			begin += 1
		} else {
			return []int{begin, end}
		}
	}
	return []int{}
}