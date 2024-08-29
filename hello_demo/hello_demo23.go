package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	left := 0
	right := len(matrix[0]) - 1
	top := 0
	bottom := len(matrix) - 1

	ans := make([]int, 0, 30)
	n := len(matrix)
	m := len(matrix[0])
	for len(ans) < n*m {
		for i := left; i <= right; i++ {
			ans = append(ans, matrix[top][i])
			fmt.Println("test1", matrix[top][i])
		}
		top++
		// if top > bottom{
		//     break
		// }

		for i := top; i <= bottom; i++ {
			ans = append(ans, matrix[i][right])
			fmt.Println("test2", matrix[i][right])
		}
		right--
		// if right < left{
		//     break
		// }
		for i := right; i >= left; i-- {
			ans = append(ans, matrix[bottom][i])
			fmt.Println("test3", top, bottom, left, right, i, matrix[bottom][i])
		}
		bottom--
		// if bottom < top{
		//     break
		// }
		for i := bottom; i >= top; i-- {
			ans = append(ans, matrix[i][left])
			fmt.Println("test4", matrix[i][left])
		}
		left++
		// if left > right{
		//     break
		// }
	}
	return ans
}
func main() {
	matrix := [][]int{
		[]int{1, 2, 3, 4},
		[]int{5, 6, 7, 8},
		[]int{9, 10, 11, 12},
	}

	ans := spiralOrder(matrix)
	fmt.Println(ans)
}
