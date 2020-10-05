package utils

func ContainsString(arr []string, str string) (int, bool) {
	for i, val := range arr {
		if val == str {
			return i, true
		}
	}
	return -1, false
}
