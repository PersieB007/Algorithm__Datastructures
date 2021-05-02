package sortingAlgo

import "fmt"

func SelectionSort() {
	array := []int{100, 10, 40, 20, 80, 15, 90}
	for i := 0; i < len(array)-1; i++ {
		var minIndex int
		minIndex = i
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[minIndex] {
				minIndex = j
			}
		}
		array[i], array[minIndex] = array[minIndex], array[i]
	}
	fmt.Println(array)
}
