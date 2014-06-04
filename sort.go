package goalgo

func insertionSort(slice []int) {
	for i := 1; i < len(slice); i++ {
		j := i
		for j > 0 && slice[j-1] > slice[j] {
			slice[j-1], slice[j] = slice[j], slice[j-1]
		}
	}
}
