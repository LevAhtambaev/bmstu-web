package api

func partition(arr []Gift, low, high int) ([]Gift, int) {
	pivot := arr[high].Priority
	i := low
	for j := low; j < high; j++ {
		if arr[j].Priority < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

func quickSort(g []Gift, low, high int) []Gift {
	if low < high {
		var p int
		g, p = partition(g, low, high)
		g = quickSort(g, low, p-1)
		g = quickSort(g, p+1, high)
	}
	return g
}
