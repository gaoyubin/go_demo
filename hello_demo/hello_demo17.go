package main

import (
	"fmt"
	"sync"
)

var fchan chan int
var bchan chan int
var wg sync.WaitGroup

func Foo() {
	defer wg.Done()
	for i := 1; i <= 10; i = i + 2 {
		<-fchan
		fmt.Println("A", i)
		bchan <- 1

	}
}
func Bar() {
	defer wg.Done()
	for i := 2; i <= 10; i = i + 2 {
		<-bchan
		fmt.Println("B", i)
		fchan <- 1
	}

}
func main() {
	// ctx, cancel := context.WithCancel(context.Background())
	fchan = make(chan int, 1)
	bchan = make(chan int)
	wg = sync.WaitGroup{}
	fchan <- 1
	wg.Add(2)

	go Foo()
	go Bar()
	wg.Wait()
	// time.Sleep(2 * time.Second)
	fmt.Println("end")

}
