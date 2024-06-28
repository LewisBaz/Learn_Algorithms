package twopointers

import "fmt"

func Run() {
	fmt.Println(RemoveDuplicates([]int{0,0,1,1,1,2,2,3,3,4}))
	fmt.Println(TwoSumHashMap([]int{0, 0, 3, 4}, 0))
	fmt.Println(TwoSumHashMap([]int{2,7,11,15}, 9))
	fmt.Println(TwoSumHashMap([]int{3,2,3}, 6))
	fmt.Println(TwoSumHashMap([]int{3,2,4}, 6))
	fmt.Println(removeElement([]int{0,1,2,2,3,0,4,2}, 2))
	fmt.Println(removeElement([]int{2,2}, 2))
}