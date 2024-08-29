package main

import "fmt"

func Partiton(arr []int, left int, right int) int {
	l := left
	r := right
	key := arr[left]
	for l < r {
		for l < r && arr[r] >= key {
			r--
		}
		for l < r && arr[l] <= key {
			l++
		}
		arr[r], arr[l] = arr[l], arr[r]
	}
	arr[left], arr[l] = arr[l], arr[left]
	return l
}
func QuickSort(arr []int, left int, right int, k int) {
	if left >= right {
		return
	}
	pos := Partiton(arr, left, right)
	if k < pos {
		QuickSort(arr, left, pos-1, k)
	} else if k > pos {
		QuickSort(arr, pos+1, right, k)
	}
}
func main() {
	arr := []int{9, 3, 4, 5, 1, 2, 6, 8}
	QuickSort(arr, 0, len(arr)-1, 4)
	fmt.Println(arr)
}
