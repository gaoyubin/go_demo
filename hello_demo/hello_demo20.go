package main

import "fmt"

func dfs(n int, m int) int {
	if n == 0 || m == 0 {
		return 0
	}

	if m == 1 {
		return n
	}

	tmp := 0
	if m < 0 {
		tmp = -1 * dfs(n, -m-1)
		// fmt.Println(tmp, n, -m-1)
		return tmp - n
	} else {
		tmp = dfs(n, m-1)
		return tmp + n
	}

	// return tmp + n
}

func dfs2(n int, m int) int {
	if n == 0 || m == 0 {
		return 0
	}

	if m == 1 {
		return n
	}
	tmp := dfs2(n, m/2)
	if m%2 == 0 {
		return tmp + tmp
	} else {
		return tmp + tmp + n
	}

}

func main() {
	// res := dfs(-3, 2)

	res := dfs(9, 3)
	fmt.Println(res)
}

//
// x=min(m,n)
// logx
