package sorting

import "fmt"

func FindTheDifference(s string, t string) byte {
	sumS := 0
	sumT := 0
	for i, v := range t {
		sumT += int(v)
		if i != len(s) {
			sumS += int(s[i])
		}
	}
	dif := rune(sumT - sumS)
	fmt.Println(string(dif))
	return byte(dif)
}