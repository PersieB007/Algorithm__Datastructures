package serchingAlgo

import "fmt"

func InterpolationSearch(array []int, value int) int {
	var first int = 0
	var last int = len(array) - 1
	var middle int = 0
	if array[last] == value {
		fmt.Println("Found at index ", last)
		return last
	}
	middle = first + (((last - first) / (array[last] - array[first])) * (value - array[first]))
	for first := 0; first <= last; {
		if array[middle] < value {
			first = middle + 1
		} else if array[middle] == value {
			fmt.Println("Found at index ", middle)
			return middle
		} else {
			last = middle - 1
		}
		middle = first + (((last - first) / (array[last] - array[first])) * (value - array[first]))
	}
	return -1
}
