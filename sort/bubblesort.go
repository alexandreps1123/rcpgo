package sort

//	(1/2)*n^2 + (1/2)*n - 1

func BubbleSort(a []int) {
	for i := 0; i < len(a)-1; i++ {
		for j := i; j < len(a); j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}
