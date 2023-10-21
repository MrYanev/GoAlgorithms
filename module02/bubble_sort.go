package module02

func BubbleSort(list []int) []int {
	var sorted = true
	lenght := len(list) - 1
	for sorted {
		sorted = false
		for i := 0; i <= lenght; i++ {
			if list[i] > list[i+1] {
				list[i], list[i+1] = list[i+1], list[i]
				sorted = true
			}
		}
		lenght--
	}

	return list
}
