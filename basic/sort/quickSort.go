package main

import "fmt"

func quickSort(arr []int, lo, hi int) {
	if lo >= hi {
		return
	}
	mid := partition(arr, lo, hi)
	quickSort(arr, lo, mid-1)
	quickSort(arr, mid+1, hi)
}

func partition(arr []int, lo, hi int) int {
	i, j, v := lo, hi, arr[lo]
	for {
		for arr[i] <= v && i < hi {
			i++
		}
		for arr[j] >= v && j > lo {
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
