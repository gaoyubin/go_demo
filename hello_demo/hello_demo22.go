package main

import "fmt"

func generateMatrix(n int) [][]int {
	top, bottom := 0, n-1
	left, right := 0, n-1
	num := n * n
	// tar := n * n
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	for num > 0 {

		for i := right; i >= left; i-- {
			matrix[bottom][i] = num
			num--
		}
		bottom--
		for i := bottom; i >= top; i-- {
			matrix[i][left] = num
			num--
		}
		left++
		for i := left; i <= right; i++ {
			matrix[top][i] = num
			num--
		}
		top++
		for i := top; i <= bottom; i++ {
			matrix[i][right] = num
			num--
		}
		right--
	}
	return matrix
}

func main() {
	ans := generateMatrix(5)
	for i := 0; i < len(ans); i++ {
		fmt.Println(ans[i])
	}
	// fmt.Println(ans)
}
