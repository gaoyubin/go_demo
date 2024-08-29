package main

type People struct {
	age  int
	name string
}

func reverse(str string) string {
	b := []byte(str)
	for i, _ := range str {
		// fmt.Println(reflect.TypeOf(str[i]))
		b[i] = b[i] + 1
	}

	return string(b)
}

func main() {
	// s := []int{1, 2, 3, 4}
	// ptr := &s[0]
	// fmt.Println(ptr)
	// // unPtr :=
	// // fmt.Println(unPtr)
	// unPtr := uintptr(unsafe.Pointer(ptr))
	// // unPtr = unPtr + 1
	// fmt.Println(unPtr, *((*int)(unsafe.Pointer(unPtr + unsafe.Sizeof(s[0])))))

	// p := &People{
	// 	age:  10,
	// 	name: "hello",
	// }
	// agePtr := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + unsafe.Offsetof(p.age)))
	// *agePtr = 12
	// fmt.Println(p)
	// str := "hello world"
	// res := reverse(str)
	// fmt.Println(res)

	// strReader := strings.NewReader("hello world\nare you ok?\ntest ok?")
	// r := bufio.NewReader(strReader)
	// for {
	// 	l, _, err := r.ReadLine()
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	fmt.Println(string(l))
	// }
	// ch1 := make(chan int)
	// ch2 := make(chan int)
	// go func() {
	// 	time.Sleep(4 * time.Second)
	// 	ch1 <- 1
	// }()
	// go func() {
	// 	time.Sleep(1 * time.Second)
	// 	<-ch2
	// }()
	// select {
	// case <-ch1:
	// 	{
	// 		fmt.Println("ch1")
	// 	}
	// case ch2 <- 1:
	// 	{
	// 		fmt.Println("ch2")
	// 	}
	// case <-time.After(3 * time.Second):
	// 	{
	// 		fmt.Println("time out")
	// 	}
	// }
	// ctx, cancel := context.WithCancel(context.Background())
	// go watch(ctx, "test1")
	// go watch(ctx, "test2")
	// go watch(ctx, "test3")
	// time.Sleep(5 * time.Second)
	// cancel()

	// s1 := []int{1, 4, 3}
	// s2 := []int{1, 3, 4}
	// if reflect.DeepEqual(s1, s2) {
	// 	fmt.Println("true")
	// }

	// a := 0
	// go func() {
	// 	for i := 0; i < 1000; i++ {
	// 		a = a + 1
	// 	}
	// }()
	//
	// go func() {
	// 	for i := 0; i < 1000; i++ {
	// 		a = a + 1
	// 	}
	// }()
	// time.Sleep(1 * time.Second)
	// fmt.Println(a)

}

// func watch(ctx context.Context, name string) {
// 	for {
// 		select {
// 		case <-ctx.Done():
// 			{
// 				fmt.Println("end")
// 				return
// 			}
// 		case <-time.After(2 * time.Second):
// 			{
// 				fmt.Println(name)
// 			}
// 		}
// 	}
//
// }

create table PeoplePo{

	identity bigint not null primary key
	age int not null
	sex int not null
	createts timestamp
	updatets timestamp
	deletets timestamp
}
// 建 age 索引

// vo
type People struct {
	identity int64
	age int
	sex int
}

type CreateReq struct{
	people People
}

type CreateRsp struct{
	people People
}

type UpdateTsReq struct{
	identity int64
	seq int
}
type UpdateTsRsp struct{
	people People
}

type QueryReq struct{
	p People
}
type QueryRsp struct{
	p People
}

func create(p people)(error,people){
	err,res := insert xx value(p.identity, p.age, p.sex)
	if err != nil{
		return err,nil
	}
	return nil, res

}
func