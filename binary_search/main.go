package binarysearch

import "fmt"

func Run() {
	// fmt.Println(Search([]int{-1,0,3,5,9,12}, 2))
	// fmt.Println(Search([]int{-1,0,3,5,9,12}, 9))
	
	fmt.Println(searchInsert([]int{1,3,5,6}, 5))
	fmt.Println(searchInsert([]int{1,3,5,6}, 2))
	fmt.Println(searchInsert([]int{1,3,5,6}, 7))
	fmt.Println(searchInsert([]int{1,3,5,6}, 0))
	fmt.Println(searchInsert([]int{1,3}, 2))
	fmt.Println(searchInsert([]int{1,3}, 0))
}