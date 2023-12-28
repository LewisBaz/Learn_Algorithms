package arrays

import (
	"fmt"
	"sort"
)

func MergeArraysInAsc(nums1 []int, m int, nums2 []int, n int) {
	i, j := 0, 0
	final := make([]int, 0, m+n)
	for i < m && j < n {
		if nums1[i] > nums2[j] {
			final = append(final, nums2[j])
			j++
		} else {
			final = append(final, nums1[i])
			i++
		}
	}
	if i != m {
		final = append(final, nums1[i:]...)
	}
	if j != n {
		final = append(final, nums2[j:]...)
	}
	nums1 = final
	fmt.Println(nums1)
}

func MergeArrays(nums1 []int, m int, nums2 []int, n int) {
	i, j := 0, 0
	nums1 = sortInAsc(nums1)
	nums2 = sortInAsc(nums2)
	nums1Copy := make([]int, m+n)
	copy(nums1Copy, nums1)
	nums1 = []int{}
	for i < m && j < n {
		if nums1Copy[i] > nums2[j] {
			nums1 = append(nums1, nums2[j])
			j++
		} else {
			nums1 = append(nums1, nums1Copy[i])
			i++
		}
	} 	
	if i != m {
		nums1 = append(nums1, nums1Copy[i:]...)
	}
	if j != n {
		nums1 = append(nums1, nums2[j:]...)
	}
	copy(nums1, nums1)
	fmt.Println(nums1)
}

func sortInAsc(nums []int) []int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	return nums
}