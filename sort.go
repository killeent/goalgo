package goalgo

func insertionSort(slice []int) {
	for i := 1; i < len(slice); i++ {
		j := i
		for j > 0 && slice[j-1] > slice[j] {
			swap(slice, j-1, j)
			j--
		}
	}
}

func swap(slice []int, i, j int) {
	temp := slice[i]
	slice[i] = slice[j]
	slice[j] = temp
}
