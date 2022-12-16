package sort

func partition(a []int, low, high int) int {
	pivot := a[high]
	i := (low - 1)

	for j := low; j <= high-1; j++ {
		if a[j] <= pivot {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}

	a[i+1], a[high] = a[high], a[i+1]

	return i + 1
}

func quickSort(a []int, s_index, l_index int) {
	if s_index < l_index {
		p_index := partition(a, s_index, l_index)
		quickSort(a, s_index, p_index-1) //	antes do pivot
		quickSort(a, p_index+1, l_index) //	depois do pivot
	}
}

func QuickSort(a []int) {
	quickSort(a, 0, len(a)-1)
}
