package main

import (
	"fmt"
	"runtime/debug"
)

// func Producer(factor int, out chan<- int) {
// 	for i := 1; ; i++ {
// 		time.Sleep(time.Second)
// 		out <- factor * i
// 	}
// }
// func Consumer(in <-chan int) {
// 	for v := range in {
// 		// time.Sleep(time.Second)
// 		fmt.Println(v)
// 	}
// }
// func main() {
// 	ch := make(chan int, 10)
// 	go Producer(3, ch)
// 	go Producer(5, ch)
// 	go Consumer(ch)
// 	time.Sleep(10 * time.Second)
// }

// func main() {
// 	var wg sync.WaitGroup
// 	for i := 0; i < 5; i++ {
// 		wg.Add(1)
// 		go func(i int, wg *sync.WaitGroup) {
// 			defer wg.Done()
// 			fmt.Printf("hello: %v\n", i)
// 		}(i, &wg)
// 	}
// 	wg.Wait()
// 	fmt.Println("close")
// }

//
// func worker(wg *sync.WaitGroup, factor int, ch chan bool) {
// 	defer wg.Done()
// 	for {
// 		time.Sleep(time.Second)
// 		select {
// 		default:
// 			fmt.Printf("hello world:%v\n", factor)
// 		case <-ch:
// 			return
// 		}
// 	}
// }
// func main() {
// 	var wg sync.WaitGroup
// 	ch := make(chan bool)
// 	for i := 0; i < 10; i++ {
// 		wg.Add(1)
// 		go worker(&wg, i, ch)
// 	}
// 	time.Sleep(5 * time.Second)
// 	close(ch)
// 	wg.Wait()
// }

// func worker(i int, ctx context.Context, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	defer func() {
// 		fmt.Printf("func end: %v\n", i)
// 	}()
// 	for {
// 		time.Sleep(time.Second)
// 		select {
// 		default:
// 			fmt.Printf("hello world:%v\n", i)
// 		case <-ctx.Done():
// 			return
// 		}
// 	}
// }
// func main() {
// 	var wg sync.WaitGroup
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	for i := 0; i < 5; i++ {
// 		wg.Add(1)
// 		go worker(i, ctx, &wg)
// 	}
// 	time.Sleep(5 * time.Second)
// 	cancel()
// 	wg.Wait()
// 	fmt.Println("close")
// }

func main() {
	fmt.Println("c")
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		fmt.Println("d")
		if err := recover(); err != nil {
			// fmt.Println(err) // 这里的err其实就是panic传入的内容
			s := string(debug.Stack())
			fmt.Printf("err=%v, stack=%s\n", err, s)
		}
		fmt.Println("e")
	}()
	// go f() // 开始调用f
	// time.Sleep(3 * time.Second)
	fmt.Println("f") // 这里开始下面代码不会再执行
}

func f() {
	fmt.Println("a")
	panic("异常信息")
	fmt.Println("b") // 这里开始下面代码不会再执行
}
