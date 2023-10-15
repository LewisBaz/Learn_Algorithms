package simple

// Находим наибольший общий делитель для a и b
func GCD(a int, b int) int {
	for b != 0 {
		// Находим остаток
		rem := a % b
		a = b
		b = rem
		// Повторяем пока не получим 0
	}
	return a
}