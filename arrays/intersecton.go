package arrays

// https://leetcode.com/problems/intersection-of-two-arrays-ii/description/

func Intersect(nums1 []int, nums2 []int) []int {
	result := make([]int, 0)
	inters := make(map[int]int,0)
	var small []int
	var big []int

	if len(nums1) > len(nums2) {
		small = nums2
		big = nums1
	} else {
		small = nums1
		big = nums2
	}

	for i := 0; i < len(small); i++ {
		inters[small[i]] += 1
	}
	for _, v := range big {
		if inters[v] != 0 {
			result = append(result, v)
			inters[v] -= 1
		}
	}

	return result
}