package main

import (
	"fmt"
	"time"
)

var fooChan chan int
var barChan chan int

func foo() {
	l := "123"
	for _, val := range l {
		<-fooChan
		fmt.Println(string(val))
		barChan <- 1
	}
}

func bar() {
	l := "abc"
	for _, val := range l {
		<-barChan
		fmt.Println(string(val))
		fooChan <- 1
	}
}
func main() {
	fooChan = make(chan int, 1)
	barChan = make(chan int)
	fooChan <- 1
	go foo()
	go bar()
	time.Sleep(2 * time.Second)
}
