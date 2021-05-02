package serchingAlgo

func BinarySearch(value int, array []int) int {

	low := 0
	high := len(array) - 1
	x := 0
	for low <= high {
		mid := (low + high) / 2
		if array[mid] < value {
			low = mid + 1
		} else {
			high = mid - 1
		}
		x = mid
	}

	if low == len(array) || array[low] != value {
		print("Number not found")
		return 0
	}
	return x
}
