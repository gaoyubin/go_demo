// package main

// import (
// 	"fmt"
// 	"sort"
// )

// var secondChan = make(chan int)
// var thirdChan = make(chan int)
// var mainChan = make(chan int)

// func first() {
// 	fmt.Println("first")
// 	secondChan <- 1
// }
// func second() {
// 	<-secondChan
// 	fmt.Println("second")
// 	thirdChan <- 1
// }
// func third() {
// 	<-thirdChan
// 	fmt.Println("third")
// 	mainChan <- 1
// }
// func main() {
// 	funcMap := map[int]func(){1: first, 2: second, 3: third}
// 	inputlist := [3]int{1, 2, 3}
// 	for _, num := range inputlist {
// 		go funcMap[num]()
// 	}
// 	<-mainChan

// 	l := []int{1, 3, 6, 8, 9, 11}
// 	sort.Ints(l)
// 	fmt.Println(l)
// 	m := make(map[string]string)
// 	m["1"] = "hello"
// 	m["2"] = "world"
// 	if val, ok := m["1"]; ok {
// 		fmt.Println(val)
// 	}
// }
