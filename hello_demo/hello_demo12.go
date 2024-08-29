package main

import "fmt"

func main() {
	arr := []int{2, 4, 6, 1, 10, 3, 5, 7, 8}
	QuickSort1(arr, 0, len(arr)-1)
	fmt.Println(arr)

}
func Partiton1(arr []int, left int, right int) int {
	l := left
	r := right
	target := arr[l]
	for l < r {
		for l < r && target <= arr[r] {
			r--
		}
		for l < r && target >= arr[l] {
			l++
		}
		arr[r], arr[l] = arr[l], arr[r]
	}
	arr[left], arr[l] = arr[l], arr[left]
	return l
}
func QuickSort1(arr []int, left int, right int) {
	if left >= right {
		return
	}
	index := Partiton1(arr, left, right)
	QuickSort1(arr, left, index-1)
	QuickSort1(arr, index+1, right)
}
