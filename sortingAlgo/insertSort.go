package sortingAlgo

import "fmt"

func InsertSort() {
	arr := []int{8, 2, 4, 3, 1}

	for i := 1; i < len(arr); i++ {
		curr := arr[i]
		prev := i - 1
		for prev >= 0 && arr[prev] > curr {
			arr[prev+1] = arr[prev]
			fmt.Println(arr)
			prev--
			arr[prev+1] = curr
		}

	}
	fmt.Println(arr)
}
