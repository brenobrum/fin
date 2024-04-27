package maxLength

// MaxLength checks if the given string lenght higher than n
func MaxLength(str string, n int) bool {
	if len(str) < n {
		return true
	}

	return false
}
