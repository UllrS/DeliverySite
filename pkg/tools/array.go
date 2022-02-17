package tools

func ReverseArray(inp []string) []string {
	var reverse []string
	for i := len(inp) - 1; i >= 0; i-- {
		reverse = append(reverse, inp[i])
	}
	return reverse
}
