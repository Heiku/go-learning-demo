package main

import "fmt"

func quickSort(arr []int, lo, hi int) {
	if lo >= hi {
		return
	}

	j := partition(arr, lo, hi)
	quickSort(arr, lo, j-1)
	quickSort(arr, j+1, hi)
}

func partition(arr []int, lo, hi int) int {
	i := lo
	j := hi
	v := arr[lo]

	for {
		for i < hi && arr[i] <= v {
			i++
		}
		for j > lo && arr[j] >= v {
			j--
		}
		if i >= j {
			break
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
	arr[lo], arr[j] = arr[j], arr[lo]
	return j
}

func main() {
	arr := []int{6, 2, 1, 7, 9, 3, 5, 4}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
