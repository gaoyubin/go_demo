package main

import (
	"fmt"
	"sync"
	"time"
)

/*
func main() {
	// var m1 map[string]string
	// m1 = make(map[string]string)
	m2 := make(map[string]string)
	m2["12"] = "21"
	for k, v := range m2 {
		fmt.Println(k, v)
	}
	fmt.Println("12")
}
*/
var counter = 0
var lock sync.Mutex

func Count() {
	lock.Lock()
	defer lock.Unlock()
	counter++

}
func main() {
	for i := 0; i < 1000; i++ {
		go Count()
	}
	time.Sleep(1500 * time.Microsecond)
	fmt.Println(counter)
}
