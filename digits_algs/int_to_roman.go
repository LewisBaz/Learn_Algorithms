package digitsalgs

import (
	"fmt"
)

func IntToRoman(num int) string {
	dict := make(map[string]string)
	dict["1"] = "I"
	dict["5"] = "V"
	dict["10"] = "X"
	dict["50"] = "L"
	dict["100"] = "C"
	dict["500"] = "D"
	dict["1000"] = "M"

	intList := intToList(num)

	resDict := make(map[int]string)

	iteration := 0

	listLen := len(intList) - 1

	for number := listLen; number >= 0; number-- {
		iterate(iteration, resDict, intList[number], number, dict)
		iteration += 1
	}

	resString := ""
	for i := 0; i <= listLen; i++ {
		resString += resDict[i]
	}
	fmt.Println(resString)
	return resString
}

func intToList(num int) []int {
	var digits []int
	for num > 0 {
		digits = append(digits, num%10)
		num /= 10
	}
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}
	return digits
}

func iterate(iter int, resDict map[int]string, n int, number int, dict map[string]string) {
	one := getIterationNum(iter, 1)
	five := getIterationNum(iter, 5)
	ten := getIterationNum(iter, 10)
	one = dict[one]
	five = dict[five]
	ten = dict[ten]
	if n == 1 {
		resDict[number] = one
		return
	}
	if n == 2 {
		resDict[number] = one + one
		return
	}
	if n == 3 {
		resDict[number] = one + one + one
		return
	}
	if n == 4 {
		resDict[number] = one + five
		return
	}
	if n == 5 {
		resDict[number] = five
		return
	}
	if n == 6 {
		resDict[number] = five + one
		return
	}
	if n == 7 {
		resDict[number] = five + one + one
		return
	}
	if n == 8 {
		resDict[number] = five + one + one + one
		return
	}
	if n == 9 {
		resDict[number] = one + ten
		return
	}
}

func getIterationNum(iter int, base int) string {
	res := fmt.Sprint(base)
	for i := 0; i < iter; i++ {
		res += "0"
	} 
	return res
}