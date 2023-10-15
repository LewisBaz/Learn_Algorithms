package simple

func Contains(arr []int, target int) bool {
	for _, val := range arr {
		if val == target {
			return true
		}
	} 
	return false
}