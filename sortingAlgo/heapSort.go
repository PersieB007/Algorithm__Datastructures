package sortingAlgo

func heapify(array []int, n, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2
	if left < n && array[largest] < array[left] {
		largest = left
	}
	if right < n && array[largest] < array[right] {
		largest = right
	}
	if largest != i {
		array[i], array[largest] = array[largest], array[i]
		heapify(array, n, largest)
	}
}

func heapSort(array []int) []int {
	n := len(array)
	//heapify
	for i := n/2 - 1; i >= 0; i-- {
		heapify(array, n, i)
	}
	//heapSort
	for i := n - 1; i >= 0; i-- {
		//swap
		array[i], array[0] = array[0], array[i]
		heapify(array, i, 0)
	}
	return array
}
