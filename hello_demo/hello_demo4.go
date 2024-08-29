package main

import "fmt"

//
//func main() {
//	ctx, cancel := context.WithCancel(context.Background())
//	go watch2(ctx)
//	go watch2(ctx)
//	time.Sleep(10 * time.Second)
//	cancel()
//	fmt.Println("结束")
//	time.Sleep(3 * time.Second)
//}
//func watch2(ctx context.Context) {
//	for {
//		select {
//		case <-ctx.Done():
//			fmt.Println("done")
//			return
//		default:
//			fmt.Println("sleep")
//			time.Sleep(2 * time.Second)
//
//		}
//	}
//}

//	func main() {
//		timer1 := time.NewTimer(3 * time.Second)
//		fmt.Println("等待三秒钟...")
//		for {
//			select {
//			case <-timer1.C:
//				fmt.Println("定时器已经触发了!")
//				timer1.Reset(time.Second)
//			}
//		}
//		fmt.Println("结束")
//	}

//func main() {
//	ch := make(chan int, 10)
//	var wg sync.WaitGroup
//	for i := 0; i < 10; i++ {
//		wg.Add(1)
//		go func() {
//			defer wg.Done()
//			time.Sleep(2 * time.Second)
//			ch <- rand.Int()
//		}()
//	}
//	wg.Wait()
//	close(ch)
//	for val := range ch {
//		fmt.Printf("val=%v\n", val)
//	}
//
//	//for {
//	//	v, ok := <-ch
//	//	if !ok {
//	//		fmt.Println("已经读取到所有数据了", v, ok)
//	//		break
//	//	}
//	//	fmt.Println("取到数据", v, ok)
//	//}
//	fmt.Println("结束")
//}

//func main() {
//	s := []int{4, 2, 3, 1}
//	sort.Ints(s)
//	fmt.Println(s)
//
//	family := []struct {
//		Name string
//		Age  int
//	}{
//		{"alice", 23},
//		{"david", 21},
//		{"eve", 2},
//		{"bob", 24},
//	}
//
//	sort.SliceStable(family, func(i, j int) bool {
//		return family[i].Age < family[j].Age
//	})
//	fmt.Println(family)
//}

//func main() {
//	l := 10
//	ch := make(chan int)
//	for i := 0; i < l; i++ {
//		go func() {
//			time.Sleep(2 * time.Second)
//			ch <- rand.Int()
//		}()
//	}
//
//	for i := 0; i < l; i++ {
//		val := <-ch
//		fmt.Println("获取到:", val)
//	}
//	fmt.Println("结束")
//}

//	func main() {
//		ch := make(chan string)
//		go func() {
//			time.Sleep(time.Second * 2)
//			ch <- "写入某个值"
//		}()
//		select {
//		case rs := <-ch:
//			fmt.Println("结果是:", rs)
//		case <-time.After(time.Second * 5):
//			fmt.Println("超时:")
//		}
//	}
func main() {
	targetI32 := 1
	targetTagOptionTotal := 2
	compareI32 := 25771869
	compareTagOptionTotal := 58475272
	if targetI32*1000 < targetTagOptionTotal*5 || compareI32*1000 < compareTagOptionTotal*5 {
		fmt.Println("yes")
	}
	fmt.Println("no")

}
