package main

import "fmt"

func main() {

	w := []int{1, 2, 5}
	v := []int{2, 5, 3}
	target := 9
	n := len(w)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, target+1)
	}
	for j := w[0]; j <= target; j++ {
		dp[0][j] = j / w[0] * v[0]
	}
	// 初始化
	fmt.Println(dp[0])
	for i := 1; i < n; i++ {
		for j := 0; j <= target; j++ {
			if j >= w[i] {
				dp[i][j] = max(dp[i-1][j], dp[i][j-w[i]]+v[i])
			} else {
				dp[i][j] = dp[i-1][j]
			}
			// dp[i][j-1] -> dp[i][j]
		}
	}
	fmt.Println(dp)

}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
