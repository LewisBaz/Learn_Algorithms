package floydstortoisehare

import "fmt"

func Run() {
	fmt.Println(FindDuplicate([]int{1,3,4,2,2}))
	fmt.Println(FindDuplicate([]int{3,1,3,4,2}))
	fmt.Println(FindDuplicate([]int{3,3,3,3,3}))
	fmt.Println(FindDuplicate([]int{4,3,1,4,2}))
	fmt.Println(FindDuplicate([]int{1,3,4,2,1}))
	fmt.Println(FindDuplicate([]int{1,1}))
	fmt.Println(FindDuplicate([]int{1,1,2}))
}