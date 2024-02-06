package digitsalgs

func MyPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	exp := abs(n)
	var num float64 = 1
	for exp > 0 {
		if exp % 2 == 1 {
			num *= x
		} 
		x *= x
		exp /= 2
	}
	if n < 0 {
		num = 1/num
	}
	return num
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}