package module02

func BubbleSort(list []int) []int {
	var sorted []int
	var srt = false
	for !srt {
		for i := 0; i <= len(list)-1; i++ {
			if list[i] > list[i+1] {
				tmp := list[i+1]
				list[i+1] = list[i]
				list[i] = tmp
			}
		}
	}

	return sorted
}
