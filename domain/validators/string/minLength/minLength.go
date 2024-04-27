package minLength

// MinLength checks if the given string with less length than n
func MinLength(str string, n int) bool {
	if len(str) < n {
		return false
	}

	return true
}
