package simple

import "math/rand"

func Randomize_array(arr []int) []int {
	maxI := len(arr) - 1
	for i := 0; i <= maxI; i++ {
		j := rand.Intn(maxI + 1)
		iVal := arr[i]
		jVal := arr[j]
		arr[i] = jVal
		arr[j] = iVal
	} 
	return arr
}


func Randomize_String_array(arr []string) []string {
	maxI := len(arr) - 1
	for i := 0; i < maxI; i++ {
		j := rand.Intn(maxI + 1)
		iVal := arr[i]
		jVal := arr[j]
		arr[i] = jVal
		arr[j] = iVal
	} 
	return arr
}