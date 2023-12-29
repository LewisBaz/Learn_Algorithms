package stringsf

func LengthOfLastWord(s string) int {
	res := []string{}
	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i]) != " " {
			res = append(res, string(s[i]))
		}
		if len(res) != 0 && string(s[i]) == " " {
			return len(res)
		} 
		if len(res) == len(s) {
			return len(res)
		}
	}
	return len(res)
}