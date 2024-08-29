package main

import (
	"fmt"

	"golang.org/x/sync/errgroup"
)

func test(arr [5]int) {
	arr[0] = 0
}
func main() {
	arr := [5]int{1, 3, 4, 5, 6}
	test(arr)
	fmt.Println(arr)
	group := errgroup.Group{}
	group.Go()
	group.Wait()

}
