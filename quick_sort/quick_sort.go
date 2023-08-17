package main

import "fmt"

func Partition(s []int, left int, right int) int {
	base_index := left
	base := s[left]
	for left < right {
		for left < right && base <= s[right] {
			right--
		}
		for left < right && base >= s[left] {
			left++
		}
		s[left], s[right] = s[right], s[left]
	}
	s[base_index], s[right] = s[right], s[base_index]
	return right
}
func QuickSort(s []int, left int, right int) {
	if left >= right {
		return
	}
	mid := Partition(s, left, right)
	QuickSort(s, left, mid-1)
	QuickSort(s, mid+1, right)
}
func main() {
	s := []int{1, 3, 6, 9, 2, 4, 5, 7, 8}
	fmt.Println(len(s))
	QuickSort(s, 0, len(s)-1)
	fmt.Println(s)
}
