package module02

func InsertionSort(list []int) []int {
	var sorted []int
	for _, i := range list {
		if len(sorted) == 0 {
			sorted = append(sorted, i)
		}
		for j := range sorted {
			if i < j {
				sorted = append(sorted[:j], append([]int{i}, sorted[j:]...)...)
			}
		}
		return append(sorted, i)

	}
	return sorted
}
