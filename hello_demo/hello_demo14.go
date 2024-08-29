package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// type Node struct {
// 	val   int
// 	left  *Node
// 	right *Node
// }
//
// func Find(root *Node) []int {
// 	if root == nil {
// 		return []int{}
// 	}
//
// 	l := []*Node{root}
// 	res := make([]int, 0, 1000)
// 	for len(l) > 0 {
// 		size := len(l)
// 		for i := 0; i < size; i++ {
// 			node := l[0]
// 			l = l[1:]
// 			if node.left != nil {
// 				l = append(l, node.left)
// 			}
// 			if node.right != nil {
// 				l = append(l, node.right)
// 			}
// 			res = append(res, node.val)
// 		}
// 	}
// 	return res
// }

// func main() {
// 	a := []int{1, 2, 3}
// 	b := []int{1, 3, 2}
// 	if reflect.DeepEqual(a, b) {
// 		fmt.Println("true")
// 	} else {
// 		fmt.Println("false")
// 	}
// 	// fmt.Println(a==b)
// 	m := make(map[int]int)
// 	fmt.Println(&a[1])
// }

func Merge(a []int, b []int) []int {
	i := 0
	j := 0
	res := make([]int, 0, len(a)+len(b))
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			res = append(res, a[i])
			i++
		} else {
			res = append(res, b[j])
			j++
		}
	}
	if i < len(a) {
		res = append(res, a[i:]...)
	}
	if j < len(b) {
		res = append(res, b[j:]...)
	}
	return res
}
func MergeSort(arr []int, left int, right int) {
	if left >= right {
		return
	}
	mid := (left + right) / 2
	MergeSort(arr, left, mid)
	MergeSort(arr, mid+1, right)
	res := Merge(arr[left:mid+1], arr[mid+1:right+1])
	copy(arr[left:right+1], res)

	// sort.Slice()
}
func main() {
	arr := []int{1, 5, 6, 2, 3, 7, 4}
	// b := []int{10, 20, 30}
	// copy(arr[4:], b)
	// fmt.Println(arr, b)
	MergeSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
	strReader := strings.NewReader("hello the world\n go echo")
	r := bufio.NewReader(strReader)
	w := []string{}
	for {
		buf, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		strSlice := strings.Split(string(buf), " ")
		w = append(w, strSlice...)
	}
	fmt.Println(w)
}
