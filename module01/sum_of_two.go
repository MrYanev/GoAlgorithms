package module01

func SumOfTwo(numList []int, sum int) (int, int) {
	for idx1, i := range numList {
		for idx2, j := range numList[idx1:] {
			if i+j == sum {
				return idx1, idx2
			}
		}
	}
	return -1, -1
}
