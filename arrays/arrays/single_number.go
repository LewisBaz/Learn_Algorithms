package arrays

// https://leetcode.com/problems/single-number/

// xor of two same number is 0
// xor is commutative
// 1^3^3^4^5^4^5=1

func singleNumber(nums []int) int {
	var result int
    for n := range nums {
		result = result ^ nums[n]
	}
	return result
}