package module01

func Fibonacci(num int) int {
	if num <= 1 {
		return 1
	} else {
		return Fibonacci(num-2) + Fibonacci(num-1)
	}
}
