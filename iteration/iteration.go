package iteration

func Repeat(s string, t int) string {
	result := ""
	for i := 0; i < t; i++ {
		result += s
	}

	return result
}
