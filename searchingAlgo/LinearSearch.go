package serchingAlgo

import "fmt"

func linearSearch(array []int, value int) {
	for _, i := range array {
		if i == value {
			fmt.Println("Found the value")
		}
	}
}
