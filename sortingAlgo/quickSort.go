package sortingAlgo

func Partition(array []int, low, high int) int {
	pivot := array[high]
	i := low - 1 //
	for j := low; j <= high-1; j++ {
		if array[j] < pivot {
			i++
			array[i], array[j] = array[j], array[i]
		}
	}
	array[i+1], array[high] = array[high], array[i+1]
	return i + 1
}

func QuickSort(array []int, low, high int) {
	if low < high {
		pi := Partition(array, low, high)
		QuickSort(array, low, pi-1)
		QuickSort(array, pi+1, high)
	}
}
