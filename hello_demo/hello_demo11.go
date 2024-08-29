package main

import (
	"fmt"
	"io"
)

// func main() {
// 	arr := []int{1, 3, 2, 4, 6}
// 	sort.Ints(arr)
// 	fmt.Println(arr)
//
// 	count := 0
// 	in := []int{}
// 	fmt.Scan(&count)
// 	for i := 0; i < count; i++ {
// 		tmp := 0
// 		fmt.Scan(&tmp)
// 		in = append(in, tmp)
// 	}
// 	fmt.Println(count, in)
// 	sort.Slice()
// }

// func main() {
// 	ch := make(chan int)
// 	go func() {
// 		time.Sleep(3 * time.Second)
// 		ch <- 1
// 	}()
//
// 	select {
// 	case <-ch:
// 		{
// 			fmt.Println("get ch")
// 		}
// 	case <-time.After(2 * time.Second):
// 		{
// 			fmt.Println("timeout")
// 		}
// 	}
//
// }
//
// func main() {
// 	strSlice := []string{"1", "32", "334", "123"}
// 	ctx, cancel := context.WithCancel(context.Background())
// 	ch := make(chan string)
// 	go func() {
// 		defer close(ch)
// 		for i, _ := range strSlice {
// 			str := strSlice[i] + "0"
// 			select {
// 			case ch <- str:
// 			case <-ctx.Done():
// 				{
// 					return
// 				}
// 			}
// 		}
// 	}()
//
// 	wg := sync.WaitGroup{}
// 	res := make([]int, 0, 200)
// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		i := 0
// 		for line := range ch {
// 			if i == 2 {
// 				cancel()
// 			}
// 			i++
// 			// fmt.Println(line)
//
// 			num, _ := strconv.Atoi(line)
// 			res = append(res, num)
// 		}
// 	}()
// 	wg.Wait()
// 	fmt.Println(res)
//
// 	errGroup := errgroup.Group{}
// 	errGroup.Go(func() error {
// 		return nil
// 	})
// 	if err := errGroup.Wait(); err != nil {
//
// 	}
// }

type Node struct {
	val  int
	next *Node
}

func GenListBySlice(in []int) *Node {
	dummy := &Node{}
	cur := dummy
	for _, val := range in {
		n := &Node{
			val: val,
		}
		cur.next = n
		cur = n
	}
	return dummy.next
}
func GetList(head *Node) []int {
	cur := head
	res := make([]int, 0, 100)
	for cur != nil {
		res = append(res, cur.val)
		cur = cur.next
	}
	return res
}
func ReverseList(head *Node) *Node {
	cur := head
	var pre *Node = nil
	for cur != nil {
		tmp := cur.next
		cur.next = pre
		pre = cur
		cur = tmp
	}
	return pre
}
func Scan() []int {
	in := make([]int, 0, 100)
	for {
		tmp := int(0)
		if _, err := fmt.Scan(&tmp); err != io.EOF {
			in = append(in, tmp)
		} else {
			break
		}
	}
	return in
}

// func main() {
// 	// count := 0
// 	// fmt.Scan(&count)
// 	in := Scan()
// 	sort.Ints(in)
// 	sort.Slice(in, func(i, j int) bool {
// 		if in[i] > in[j] {
// 			return true
// 		} else {
// 			return false
// 		}
// 	})
// 	fmt.Println(in)
// 	head := GenListBySlice(in)
// 	reverseHead := ReverseList(head)
// 	res := GetList(reverseHead)
// 	fmt.Println(res)
// 	strconv.Atoi()
// }

func main() {
	ans := [][]int{}
	tmp := []int{1, 2, 3}
	ans = append(ans, tmp)
	fmt.Println(ans)
	tmp[0], tmp[1] = tmp[1], tmp[0]
	fmt.Println(ans)
	strconv.
}
