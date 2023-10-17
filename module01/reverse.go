package module01

func Reverse(word string) string {
	var reversed string
	for i := len(word) - 1; i >= 0; i-- {
		reversed = reversed + string(word[i])
	}
	return reversed
}
