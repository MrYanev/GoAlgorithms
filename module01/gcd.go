package module01

func Gcd(a, b int) int {
	var ans int
	var factA, factB, ansL []int
	for i := 0; i <= a; i++ {
		if a%i == 0 {
			factA = append(factA, i)
		}
	}
	for j := 0; j <= a; j++ {
		if a%j == 0 {
			factB = append(factB, j)
		}
	}
	for _, numA := range factA {
		for _, numB := range factB {
			if numA == numB {
				ansL = append(ansL, numA)
			}
		}
	}
	for sum := range ansL {
		ans += sum
	}
	return ans
}
