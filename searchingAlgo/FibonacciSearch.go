package serchingAlgo

var (
	fiboSeries []int
)

func fibo(n int) int {
	a := 0
	b := 1
	if n < 0 {
		fiboSeries = append(fiboSeries, 0)
		return 0
	} else if n == 1 {
		fiboSeries = append(fiboSeries, 1)
		return 1
	} else {
		for i := 1; i < n; i++ {
			c := a + b
			a = b
			b = c
			fiboSeries = append(fiboSeries, b)
		}
		return b
	}

}

func helper(array []int) int {
	n := len(array)
	return BinarySearch(n, fiboSeries)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func fibonnaciSearch(array []int, value int) int {
	fibo(len(array))
	index := helper(array)
	// fibM := fiboSeries[index+1]
	fibM1 := fiboSeries[index]
	fibM2 := fiboSeries[index-1]
	offset := 0
	n := len(array)
	for fibM1 > 1 {
		i := min(offset+fibM2, n-1)
		if array[i] < value {
			fibM1 = fiboSeries[index-1]
			fibM2 = fiboSeries[index-2]
			offset = i
		} else if array[i] > value {
			fibM1 = fiboSeries[index-2]
			fibM2 = fiboSeries[index-3]
		} else {
			return array[i]
		}
	}

	return -1
}
