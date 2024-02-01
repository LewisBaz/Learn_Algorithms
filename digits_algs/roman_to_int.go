package digitsalgs

import (
	"fmt"
	"strings"
)

func RomanToInt(s string) int {
	dict := make(map[string]int)
	dict["I"] = 1
	dict["V"] = 5
	dict["X"] = 10
	dict["L"] = 50
	dict["C"] = 100
	dict["D"] = 500
	dict["M"] = 1000

	arr := strings.Split(s, "")
	reversed := make([]string, len(arr))
    for i := 0; i < len(arr); i++ {
        reversed[i] = arr[len(arr)-1-i]
    }

	res := make([]int, len(reversed))

	counter := len(reversed) - 1
    for index, char := range reversed {
		var prev string
		if index - 1 >= 0 {
			prev = reversed[index - 1]	
		}
		if char == "I" && (prev == "V" || prev == "X") {
			res[counter + 1] = dict[prev] - 1
			counter -= 1
			continue
		}
		if char == "X" && (prev == "L" || prev == "C") {
			res[counter + 1] = dict[prev] - 10
			counter -= 1
			continue
		}
		if char == "C" && (prev == "D" || prev == "M") {
			res[counter + 1] = dict[prev] - 100
			counter -= 1
			continue
		}
			
		res[counter] = dict[char]
		counter -= 1
	}

	var num int
	for _, n := range res {
		num += n
	}

	fmt.Println(num)
	return num
}