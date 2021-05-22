package sortingAlgo

func Merge(n int, left, right []int) []int {

	i, j := 0, 0
	sorted_array := make([]int, n)

	for k := 0; k < n; k++ {
		if i > len(left)-1 && j <= len(right)-1 {
			sorted_array[k] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			sorted_array[k] = left[i]
			i++
		} else if left[i] < right[j] {
			sorted_array[k] = left[i]
			i++
		} else {
			sorted_array[k] = right[j]
			j++
		}
	}
	return sorted_array
}

func MergeSort(array []int) []int {
	n := len(array)
	if len(array) < 2 {
		return array
	}
	mid := (len(array)) / 2
	return Merge(n, MergeSort(array[:mid]), MergeSort(array[mid:]))
}
