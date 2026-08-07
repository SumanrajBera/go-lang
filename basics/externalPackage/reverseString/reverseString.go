package reversestring

func Reverse(str string) string {
	var output string = ""

	for i := len(str) - 1; i >= 0; i-- {
		output = output + string(str[i])
	}

	return output
}
