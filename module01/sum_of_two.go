package module01

func SumOfTwo(numList []int, sum int) (int, int) {
	var ans1, ans2 int

	for idx1, i := range numList {
		for idx2, j := range numList[idx1:] {
			if i+j == sum {
				ans1 = idx1
				ans2 = idx2
			}
		}
	}
	return ans1, ans2
}
