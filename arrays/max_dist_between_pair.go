package arrays

import (
	"fmt"
)

// NOT COMPLETED

// You are given two non-increasing 0-indexed integer arrays nums1​​​​​​ and nums2​​​​​​.

// A pair of indices (i, j), where 0 <= i < nums1.length and 0 <= j < nums2.length, is valid if both i <= j and nums1[i] <= nums2[j]. The distance of the pair is j - i​​​​.

// Return the maximum distance of any valid pair (i, j). If there are no valid pairs, return 0.

// An array arr is non-increasing if arr[i-1] >= arr[i] for every 1 <= i < arr.length.

func MaxDistance(nums1 []int, nums2 []int) int {
	var maxDistance int

	valid := [][]int{}

	n1 := 0
	n2 := len(nums2) - 1

	counter := 0

	for n1 <= n2 {
		counter += 1
		if len(nums1) <= n1 || len(nums2) <= n2 {
			fmt.Println(counter)
			fmt.Println(valid)
			return maxDistance
		}

		if nums1[n1] <= nums2[n2] {
			valid = append(valid, []int{n1, n2})
			diff := n2 - n1
			if diff > maxDistance {
				maxDistance = diff
				// if n2 == len(nums2) - 1 {
				// 	fmt.Println(counter)
				// 	return maxDistance
				// }
			}
		}
		if n1 == n2 || nums2[n2-1] == nums2[n2] {
			n1 += 1
			n2 = len(nums2) - 1
		} else {
			n2 -= 1
		}
	}

	fmt.Println(valid)
	fmt.Println(counter)

	return maxDistance
}

func MaxDistanceEasy(nums1 []int, nums2 []int) int {
	var maxDistance int

	valid := [][]int{}

	for i, n1 := range nums1 {
		for j, n2 := range nums2 {
			if n1 <= n2 && i <= j {
				valid = append(valid, []int{i, j})
				if j - i > maxDistance {
					maxDistance = j - i
				}
			}
		}
	}
	fmt.Println(valid)

	return maxDistance
}
