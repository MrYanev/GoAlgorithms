package module01

func Factor(primes []int, number int) []int {
	var ans []int
	for _, num := range primes {
		for number%num == 0 {
			ans = append(ans, number)
			number = number / num
		}
	}
	if number > 1 {
		ans = append(ans, number)
	}
	return ans
}
