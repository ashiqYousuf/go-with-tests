package iteration

func Repeat(char string, k int) string {
	var repeated string
	for i := 1; i <= k; i++ {
		repeated += char
	}
	return repeated
}

func RepeatOptimized(char string, k int) string {
	if char == "" {
		return ""
	}

	var repeated string
	for i := 1; i <= k; i++ {
		repeated += char
	}
	return repeated
}
