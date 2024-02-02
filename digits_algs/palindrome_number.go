package digitsalgs

import (
	"strconv"
)

func IsPalindrome(x int) bool {
	stringX := strconv.Itoa(x)
    length := len(stringX)
	var result bool = true
	// if length == 1 {
	// 	return true
	// }
	for i := 0; i < length; i++ {
		firstIndex := i
		lastIndex := length - firstIndex - 1
		if firstIndex != lastIndex {
			if stringX[firstIndex] == stringX[lastIndex] {
				result = true
			} else {
				result = false
				return result
			}
		} else {
			return result
		}
	}
	return result
}