package stringsf

func IsIsomorphic(s string, t string) bool {
	container := make(map[rune]rune)
	for i, char := range s {
		if container[char] == 0 {
			for _, v := range container {
				if v == rune(t[i]) {
					return false
				}
			}
			container[char] = rune(t[i])
		} else if container[char] != rune(t[i]) {
			return false
		}
	}
	return true
}
