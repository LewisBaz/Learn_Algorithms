package arrays

import "fmt"

func MajorityElement(nums []int) int {
	var element int
	counter := 0
	for _, v := range nums {
		if counter == 0 {
			element = v
			counter += 1
			continue
		} else {
			if element == v {
				counter += 1
			} else {
				counter -= 1
			}
		}

	}
	fmt.Println(element)
	return element
}
