package isNotEmpty

// IsNotEmpty validates if a string is empty
func IsNotEmpty(str string) bool {
	if len(str) != 0 {
		return true
	}

	return false
}
