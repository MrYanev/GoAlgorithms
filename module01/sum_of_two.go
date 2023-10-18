package module01

func SumOfTwo(numList []int, sum int) []int {
	var ans []int
	for _, i := range numList {
		for _, j := range numList[i:] {
			if i+j == sum {
				ans = append(ans, i)
				ans = append(ans, j)
			}
		}
	}
	return ans
}
