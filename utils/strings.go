package utils

func Occur(slice []string, value string) int {
	for i, x := range slice {
		if x == value {
			return i
		}
	}
	return -1
}

func Startwith(x string, a string) bool {
	if x[:len(a)] == a {
		return true
	}
	return false
}
