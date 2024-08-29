package main

import "fmt"

func Create(n int) [][]int {
	// i := n / 2
	// j := n / 2
	cnt := 1
	ans := make([][]int, n)
	for i := 0; i < n; i++ {
		ans[i] = make([]int, n)
	}
	left := n / 2
	right := n / 2
	top := n / 2
	button := n / 2
	for {
		right++
		if right >= n {
			break
		}
		for i := left; i <= right; i++ {
			ans[button][i] = cnt
			fmt.Println("test1", button, i, ans[button][i])
			cnt++
		}
		// fmt.Println("test1", ans[button][right], cnt)
		top--
		if top < 0 {
			break
		}
		for i := button + 1; i >= top; i-- {
			ans[i][right] = cnt
			fmt.Println("test2", i, right, ans[i][right])
			cnt++
		}

		left--
		if left < 0 {
			break
		}
		for i := right - 1; i >= left; i-- {
			ans[top][i] = cnt
			fmt.Println("test3", top, i, ans[top][i])
			cnt++
		}

		button++
		if button >= n {
			break
		}
		for i := top + 1; i >= button; i++ {
			ans[i][left] = cnt
			fmt.Println("test4", i, left, ans[i][left])
			cnt++
		}

		//
		// right++
		// if right > n {
		// 	break
		// }

	}
	return ans
}

func main() {
	ans := Create(5)
	for i := 0; i < len(ans); i++ {
		fmt.Println(ans[i])
	}
	// fmt.Println(ans)
}
