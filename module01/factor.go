package module01

func Factor(primes []int, number int) []int {
	var ans []int
	for _, num := range primes {
		if number%num == 0 {
			number = number / num
			ans = append(ans, num)
		}
	}
	return ans
}
