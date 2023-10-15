package digitsalgs

import (
	"fmt"
	"main/simple"
	"math/rand"
)

func EqualValuesDice() int {
	vals := []int{}
	vals = append(vals, diceThrower())

	totalIterations := 0

	for i := 0; i < 6; i++ {
		next := diceThrower()
		if simple.Contains(vals, next) {
			i = 0
			vals = []int{}
			vals = append(vals, diceThrower())
			totalIterations += 1
			continue
		}
		vals = append(vals, next)
	}
	fmt.Println(vals)
	fmt.Printf("total iterations %d", totalIterations)
	return vals[0]
}

func diceThrower() int {
	val := rand.Intn(6) + 1
	return val
}