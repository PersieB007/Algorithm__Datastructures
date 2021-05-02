package sortingAlgo

import "fmt"

func BubbleSort() {
	x := []int{100, 10, 40, 20, 80, 15, 90}
	for i := 0; i < len(x)-1; i++ {
		for j := i + 1; j < len(x); j++ {
			if x[i] > x[j] {
				x[i], x[j] = x[j], x[i]
			} else {
				continue
			}
		}
	}
	fmt.Println(x)
}
