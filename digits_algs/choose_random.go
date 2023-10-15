package digitsalgs

import "math/rand"

func ChooseRandomFrom(arr []int, count int) []int {
	maxI := len(arr)
	for i := 0; i < count; i++ {
		j := rand.Intn(maxI)
		iVal := arr[i]
		jVal := arr[j]
		arr[i] = jVal
		arr[j] = iVal
	} 
	return arr[:count]
}