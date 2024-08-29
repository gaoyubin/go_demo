package main

import (
	"fmt"
	"sync"
)

/*
var a chan int
var b chan int

	func PrintA() {
		for {
			<-a
			fmt.Println("a")
			b <- 1
		}
	}

	func PrintB() {
		for {
			<-b
			fmt.Println("b")
			a <- 1
		}
	}

	func main() {
		a = make(chan int, 1)
		b = make(chan int)
		a <- 1
		go PrintA()
		go PrintB()
		time.Sleep(1 * time.Second)
	}
*/
var a chan int
var b chan int

func PrintA() {
	fmt.Println("a")
	a <- 1

}

func PrintB() {
	<-a
	fmt.Println("b")
	b <- 1

}

func PrintC() {
	<-b
	fmt.Println("c")
}

//	func main() {
//		a = make(chan int)
//		b = make(chan int)
//		go PrintA()
//		go PrintB()
//		go PrintC()
//		time.Sleep(1 * time.Second)
//	}
func testPanic() {
	panic("test panic")
	fmt.Println("test end")
}
func test() {
	defer func() {
		fmt.Println("defer begin")
		if err := recover(); err != nil {
			fmt.Println("recover:", err)
		}
	}()
	// panic("异常")
	// s := []int{}
	// fmt.Println(s[1])
	testPanic()
	fmt.Println("end")
}
func main() {
	test()
	fmt.Println("main end")
}
func do(mutex *sync.Mutex) {
	mutex.Lock()
	if a == 1 {
		mutex.Unlock()
		wg.wait()
		return 拿缓存
	}
	// 打个标记
	// a = 1
	// wg.add(1)
	// mutex.unlock
	// 获取缓存
	// wg.done

}
