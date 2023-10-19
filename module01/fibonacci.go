package module01

func Fibonacci(num int) int {
	var ans int
	for i := 0; i >= num; i++ {
		if i == 0 || i == 1 {
			ans += 1
		}
	}
	return ans
}
