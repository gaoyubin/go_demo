package main

import (
	"fmt"
	"log"
	"reflect"
	"time"
)

/*
type Person struct {
	Name string
	Age  int
}

func main() {

	//sync.Map
	// http.Handle("/tmpfiles/", http.StripPrefix("/tmpfiles/", http.FileServer(http.Dir("tmp"))))

	// http.ListenAndServe(":8888", nil)

	// http.HandleFunc()

	//p := Person{"longshuai", 12}
	//tmpl, _ := template.New("test").Parse("Name: {{.Name}}, Age: {{.Age}}")
	//tmpl, _ := template.New("test").Parse("")
	//_ = tmpl.Execute(os.Stdout, p)
	s := []int{1, 4, 5}
	str := `{{ range . }}
	sss:	{{- . -}}
{{end}}`
	tmpl, _ := template.New("test").Parse(str)
	tmpl.Execute(os.Stdout, s)

	// sync.Map
	// rate.NewLimiter()
	// rate.NewLimiter()
	rate.NewLimiter(10, 20)
	// rate.NewLimiter
	// rate.NewLimiter
	//ratelimit.New

}
*/

func producer(in chan<- int) {
	for i := 0; i < 10; i++ {
		time.Sleep(2 * time.Second)
		data := i * i
		fmt.Println("生产者生产数据:", data)
		in <- data
	}
}

func consumer(out <-chan int) {
	for data := range out {
		fmt.Println("消费者得到数据", data)
	}
}

/*
func main() {
	// ch := make(chan int)
	// go producer(ch)
	// go consumer(ch)
	// fmt.Scanln()

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			time.Sleep(1 * time.Second)
			c1 <- "one"
		}
	}()
	go func() {
		for {
			time.Sleep(2 * time.Second)
			c2 <- "two"
		}
	}()
	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println("receive", msg1)
			case msg2 := <-c2:
				fmt.Println("receive", msg2)
			}
		}
	}()
	fmt.Scanln()
}
*/

/*
type Task struct {
	id      int
	randnum int
}
type Result struct {
	task   Task
	result int
}

var tasks = make(chan Task, 10)
var results = make(chan Result, 10)
var done = make(chan bool)

func process(task Task) Result {
	sum := 0
	num := task.randnum
	for num != 0 {
		digit := num % 10
		sum += digit
		num /= 10
	}
	time.Sleep(2 * time.Second)
	return Result{task, sum}
}
func worker() {
	for task := range tasks {
		result := process(task)
		results <- result
	}
}
func getResult() {
	for result := range results {
		fmt.Println(result)
	}
	done <- true
}
func createWorkerPool(wokerNum int) {
	for i := 0; i < wokerNum; i++ {
		go worker()
	}
}
func allocate(taskNum int) {
	for i := 0; i < taskNum; i++ {
		randnum := rand.Intn(999)
		task := Task{i, randnum}
		tasks <- task
	}

}
func main() {
	go createWorkerPool(10)
	go allocate(100)
	go getResult()
	<-done
}

/*
func main() {
	var wg sync.WaitGroup
	urls := []string{
		"http://www.github.com/",
		"https://www.qiniu.com/",
		"https://www.golangtc.com/"}
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			_, err := http.Get(url)
			fmt.Println(url, err)
		}(url)
	}
	wg.Wait()
	fmt.Println("over")
}
*/

/*
type Person struct {
	Name string
	Age  int
}
type ByAge []Person

// func (a ByAge) Len(int){return len(a)}

func main() {
	s := []int{4, 2, 3, 1}
	sort.Ints(s)
	fmt.Println(s)

	family := []struct {
		Name string
		Age  int
	}{
		{"alice", 23},
		{"david", 2},
		{"eve", 25},
	}
	sort.Slice(family, func(i, j int) bool {
		return family[i].Age < family[j].Age
	})
	fmt.Println(family)


		// // num_int := 123
		// num_str := "123"
		// // num_int, err := strconv.Atoi(num_str)
		// num_int, err := strconv.ParseInt(num_str, 16, 64)
		// fmt.Println(num_int, err)

	num_int := 123
	num_str := strconv.Itoa(num_int)
	fmt.Println(num_str)

	l := list.New()
	l.PushBack("canon")
	l.PushBack(67)
	for elem := l.Front(); elem != nil; elem = elem.Next() {
		fmt.Println(elem.Value)
	}

	var x interface{}
	x = "hello"
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case float64:
		fmt.Println("float")
	default:
		fmt.Println("unknow")
	}
}
*/

/*
func testFunc(arr [3]int, slice []int) {
	arr[0] = arr[0] * 100
	slice[0] = slice[0] * 100
	// fmt.Println(&arr, &slice)
	fmt.Printf("%p %p %p\n", &arr, &slice, &slice[0])
}

func main() {
	myArr := [3]int{1, 3, 5}
	mySlice := []int{1, 3, 5}
	fmt.Printf("%p %p %p\n", &myArr, &mySlice, &mySlice[0])
	testFunc(myArr, mySlice)
	fmt.Println(myArr[0], mySlice[0])
}
*/
// func main() {
// 	a := []int{1, 3, 5, 7, 9}
// 	b := make([]int, 0, len(a))
// 	copy(b, a)
// 	fmt.Println(len(b))
// }

// func main() {
// a := make([]byte, 5, 5)
// fmt.Printf("%p\n", &a[0])
// x := append(a, 'a')
// fmt.Printf("%p %p\n", &a[0], &x[0])
// fmt.Println(cap(a), cap(x))

// fmt.Println(x, a, len(a))
// fmt.Printf("%p %p\n", &a1, &a)

// a2 := append(a, 'b')
// fmt.Println(string(a1), string(a2), a)

// nums := []int{1, 2, 3}
// result := append(nums[:], 4)
// fmt.Println(nums, result)

// s := []int{1, 3, 5, 7, 9, 11}
// s1 := s[2:]
// s2 := s[4:6]
// s[2] = 12
// fmt.Println(s, s1, s2)
// }

// func main() {
// 	var a = make([]int, 5, 10)
// 	fmt.Println(a)
// 	fmt.Printf("%p\n", a)
// 	for i := 0; i < 10; i++ {
// 		a = append(a, i)
// 		//%p 打印切片地址
// 		fmt.Printf("%v,%p,cap(a):%d,len(a):%d\n", a, a, cap(a), len(a))
// 	}
// }

// func main() {
// 	a := make([]byte, 0, 1)
// 	a1 := append(a, 'a')
// 	a2 := append(a, 'b')
// 	fmt.Println(string(a1) == string(a2))
// 	fmt.Printf("%p %p\n", &a1[0], &a2[0])
// }

// func main() {
// 	var a uint = 1
// 	var b uint = 2
// 	fmt.Println(a - b)
// }

// func main() {
// 	fmt.Println(0.1+0.2 == 0.3)
// 	a := 0.1
// 	b := 0.2
// 	fmt.Println(a + b)
// }

// func funcA(n int) func() {
// 	n++
// 	return func() {
// 		fmt.Println(n)
// 	}
// }

// func funcB(n int) func() {
// 	return func() {
// 		n++
// 		fmt.Println(n)
// 	}
// }

// func main() {
// 	f1 := funcA(10)
// 	f1()
// 	f1()

// 	f2 := funcB(10)
// 	f2()
// 	f2()
// }

// func main() {
// 	a := make([]byte, 1, 1)
// 	fmt.Printf("%p\n", &a[0])
// 	a1 := append(a, 'a')
// 	// a2 := append(a, 'b')
// 	// fmt.Println(string(a1) == string(a2))
// 	fmt.Printf("%p %p\n", &a1[0], &a[0])
// }

// func func1() (i int) {
// 	defer func() {
// 		i++
// 	}()

// 	return i
// }

// func func2() int {
// 	var i int
// 	defer func() {
// 		i++
// 	}()

// 	return i
// }

// func main() {
// 	fmt.Println(func1(), func2())
// }

// func main() {
// 	t := struct {
// 		time.Time
// 		N int
// 	}{
// 		Time: time.Date(2021, 3, 30, 0, 0, 0, 0, time.UTC),
// 		N:    5,
// 		// time.Date(2021, 3, 30, 0, 0, 0, 0, time.UTC),
// 	}
// 	fmt.Println(t.N)
// 	m, _ := json.Marshal(t)
// 	fmt.Printf("%s", m)
// }

// type Dog struct {
// 	Name string
// 	Age  int
// }
// type User struct {
// 	Name string
// 	Age  int
// 	Dog  *Dog
// }

// func (d *Dog) String() string {
// 	return "{\"name" + "\": \"" + d.Name + "\"," + "\"" + "age\": \"" + strconv.Itoa(d.Age) + "\"}"
// }
// func (u *User) String() string {
// 	return "{\"name" + "\": \"" + u.Name + "\", \"" + "age\": \"" + strconv.Itoa(u.Age) + "\", \"dog\": " + u.Dog.String() + "}"
// }

// func main() {
// 	dog := Dog{
// 		Name: "旺财",
// 		Age:  2,
// 	}
// 	user := User{
// 		Name: "张三",
// 		Age:  95,
// 		Dog:  &dog,
// 	}
// 	// fmt.Printf("%v\n", user)
// 	// fmt.Printf("%+v\n", user)
// 	// fmt.Printf("%#v\n", user)

// 	// fmt.Println(user)
// 	// fmt.Println(user.Age, user.Dog.Age)

// 	byteUser, _ := json.Marshal(&user)
// 	// fmt.Println(string(byteUser))
// 	fmt.Println(string(byteUser))

// 	var str string = "test"
// 	var data []byte = []byte(str)
// 	fmt.Println(data)
// 	var ss string = string(data)
// 	fmt.Println(ss)

// 	// strings.Count()
// 	x := []int{1,4,2,3,5,7,6,9}
// 	sort.Ints(x)
// 	fmt.Println(x)
// }

// func main() {
// 	a1 := []int{10}
// 	a2 := a1[1:]
// 	fmt.Println(a2)
// }
// type Person struct {
// 	Name    string
// 	Age     int
// 	Gender  string
// 	address string
// }

// func main() {
// 	// var i = 3.14159
// 	// v := reflect.ValueOf(&i)
// 	// e := v.Elem()
// 	// t := reflect.TypeOf(i)
// 	// e.SetFloat(1.2)
// 	// fmt.Println(e, t.Kind(), t)

// 	var p = Person{"low", 30, "男", "广州"}
// 	var v = reflect.ValueOf(&p)
// 	var e = v.Elem()
// 	var t = reflect.TypeOf(p)
// 	fmt.Println(v, t, t.Kind())

// 	var vOfName = e.FieldByName("Name")
// 	ok := vOfName.CanSet()
// 	fmt.Println(ok)
// 	vOfName.SetString("banana")
// 	fmt.Println(vOfName)
// 	var num = e.NumField()
// 	for i := 0; i < num; i++ {
// 		var val = e.Field(i)
// 		if val.Type().Kind() == reflect.String {
// 			// e := val.Elem()
// 			// ok := e.CanSet()
// 			// fmt.Println(ok)
// 			if ok := val.CanSet(); ok {
// 				// val.SetString(val.String() + "修改")
// 			}
// 			// fmt.Println(ok, val)
// 		}
// 		// fmt.Println(i, val.Type(), val)
// 	}
// 	// fmt.Println(num)
// 	fmt.Println(p)

// 	var a int
// 	typeOfA := reflect.TypeOf(a)
// 	aIns := reflect.New(typeOfA)
// 	fmt.Println(aIns.Type())
// 	inter := aIns.Interface()
// 	vv := inter.(*int)
// 	fmt.Println(vv)
// 	var x, y float64 = 1, 2
// 	z := math.Max(x, y)
// 	fmt.Println(z)
// }

// type Gender int

// const (
// 	Male Gender = iota
// 	Female
// )

// const (
// 	x1 string = "男性"
// 	x2 string = "女性"
// )

// func (g *Gender) String() string {
// 	switch *g {
// 	case Male:
// 		return "男性"
// 	case Female:
// 		return "女性"
// 	default:
// 		return ""
// 	}
// }
// func main() {
// 	g := Male
// 	fmt.Println(&g)

// 	x := x1
// 	fmt.Println(x)
// }

// func add(a, b int) int {
// 	return a + b
// }
// func main() {
// 	funcValue := reflect.ValueOf(add)
// 	paramlist := []reflect.Value{reflect.ValueOf(10), reflect.ValueOf(20)}
// 	retlist := funcValue.Call(paramlist)
// 	// fmt.Println(retlist[0])
// 	inter := retlist[0].Interface()
// 	v := inter.(int)
// 	fmt.Println(v)
// }
// func main() {
// 	// data := []byte("go语言入门教程")
// 	// rd := bytes.NewReader(data)
// 	// bufio.NewReader(rd)

// 	// data := []byte("hello world")
// 	// reader := bytes.NewReader(data)
// 	// reader := strings.NewReader("clear is better than clever")
// 	// buf := make([]byte, 4)
// 	// if _, err := reader.Read(buf); err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	// fmt.Println(string(buf))

// 	var buf bytes.Buffer
// 	buf.Write([]byte("hello world, "))
// 	fmt.Println(buf)
// 	fmt.Println(string(buf.Bytes()))

// 	// filepath := "./output.txt"
// 	// file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0666)
// 	// if err != nil {
// 	// 	fmt.Println("打开文件失败", err)
// 	// 	return
// 	// }
// 	// defer file.Close()
// 	// // bufio.NewWriter(file)
// 	// // file.Write()
// 	// writer := bufio.NewWriter(file)
// 	// str := "http://c.biancheng.net/golang/\n"
// 	// for i := 0; i < 3; i++ {
// 	// 	// writer.WriteString(str)
// 	// 	writer.Write([]byte(str))
// 	// }
// 	// writer.Flush()

// 	filepath := "./output.txt"
// 	file, err := os.Open(filepath)
// 	if err != nil {
// 		fmt.Println("文件打开失败", err)
// 	}
// 	reader := bufio.NewReader(file)
// 	for {
// 		str, err := reader.ReadString('\n')
// 		if err == io.EOF {
// 			break
// 		}
// 		fmt.Print(str)
// 	}
// 	fmt.Println("文件读取结束")

// }

// const (
// 	HintHeaderSize = 20
// )

// func encodeHint(tStamp, kSz, valueSz uint32, valuePos uint64, key []byte) []byte {
// 	buf := make([]byte, HintHeaderSize+len(key))
// 	binary.LittleEndian.PutUint32(buf[0:4], tStamp)
// 	binary.LittleEndian.PutUint32(buf[4:8], kSz)
// 	binary.LittleEndian.PutUint32(buf[8:12], valueSz)
// 	binary.LittleEndian.PutUint64(buf[12:HintHeaderSize], valuePos)
// 	copy(buf[HintHeaderSize:], []byte(key))
// 	return buf
// }
// func main() {
// 	dirFp, err := os.OpenFile(".", os.O_RDONLY, os.ModeDir)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	lists, err := dirFp.Readdirnames(-1)
// 	fmt.Println(lists)

// 	k := "hello"

// 	buf := encodeHint(uint32(time.Now().Unix()), uint32(len(k)), 12, 13, []byte(k))
// 	fmt.Println(buf)
// 	tStamp := binary.LittleEndian.Uint32(buf[0:4])
// 	kSz := binary.LittleEndian.Uint32(buf[4:8])
// 	valueSz := binary.LittleEndian.Uint32(buf[8:12])
// 	valuePos := binary.LittleEndian.Uint64(buf[12:HintHeaderSize])
// 	key := buf[HintHeaderSize : HintHeaderSize+kSz]
// 	fmt.Println(tStamp, kSz, valueSz, valuePos, string(key))

// 	// file := "/23.hint"
// 	// s := strings.LastIndex(file, "/") + 1
// 	// e := strings.LastIndex(file, ".hint")
// 	// fileID, _ := strconv.ParseInt(file[s:e], 10, 32)
// 	// fmt.Println(fileID, s, e)

// 	url := "https://developer.mozilla.org/zh-CN/search?q=URL&o=123"
// 	s := strings.Index(url, "//") + 2
// 	e := strings.Index(url[s:], "/")
// 	domain := url[s : s+e]
// 	protocal_end := strings.Index(url, ":")
// 	protocal := url[:protocal_end]
// 	fmt.Println(s, e, domain, protocal, protocal_end)

// 	var_s := strings.Index(url, "?") + 1
// 	var_var := url[var_s:]
// 	fmt.Println(var_var)

// 	var_list := strings.Split(var_var, "&")
// 	fmt.Println(var_list)
// 	var_map := make(map[string]string)
// 	for _, val := range var_list {
// 		tmp_list := strings.Split(val, "=")
// 		if len(tmp_list) != 2 {
// 			fmt.Println("var err", val, tmp_list)
// 			continue
// 		}
// 		var_map[tmp_list[0]] = tmp_list[1]
// 	}
// 	fmt.Println(var_map)

// 	tmp := "1234"
// 	tmp_int, err := strconv.ParseInt(tmp, 10, 32)
// 	tmp_str := strconv.Itoa(int(tmp_int))
// 	fmt.Println(tmp_int, tmp_str)
// }

// type User struct {
// 	Name string
// 	Age  int
// }

// func main() {
// 	user1 := &User{
// 		Name: "low",
// 		Age:  12,
// 	}
// 	result, _ := json.Marshal(user1)
// 	fmt.Println(string(result))

// 	var user2 User
// 	jsonString := `{"Name": "kuizuo", "Age" : 20}`
// 	json.Unmarshal([]byte(jsonString), &user2)
// 	fmt.Println(user2)

// 	// json.NewEncoder()
// 	var user3 User
// 	data := []byte("{\"Name\": \"kyun\", \"Age\" : 25}")
// 	rd := bytes.NewReader(data)
// 	bufrd := bufio.NewReader(rd)
// 	json.NewDecoder(bufrd).Decode(&user3)
// 	fmt.Println(user3)

// 	var buf bytes.Buffer
// 	json.NewEncoder(&buf).Encode(user1)
// 	fmt.Println("show encoder", string(buf.Bytes()))

// 	var buf2 bytes.Buffer
// 	gob.NewEncoder(&buf2).Encode(user1)
// 	fmt.Println("show gob", string(buf2.Bytes()))

// 	var user4 User
// 	gob.NewDecoder(&buf2).Decode(&user4)
// 	fmt.Println("show gob decode", user4)
// }

// func main() {
// 	qcrao := Student{age: 18}
// 	whatJob(&qcrao)

// 	growUp(&qcrao)
// 	fmt.Println(qcrao)

// 	stefno := Programmer{age: 100}
// 	whatJob(&stefno)

// 	growUp(&stefno)
// 	fmt.Println(stefno)
// }

// func whatJob(p Person) {
// 	p.job()
// }

// func growUp(p Person) {
// 	p.growUp()
// }

// type Person interface {
// 	job()
// 	growUp()
// }

// type Student struct {
// 	age int
// }

// func (p *Student) job() {
// 	fmt.Println("I am a student.")
// 	return
// }

// func (p *Student) growUp() {
// 	p.age += 1
// 	return
// }

// type Programmer struct {
// 	age int
// }

// func (p Programmer) job() {
// 	fmt.Println("I am a programmer.")
// 	return
// }

// func (p Programmer) growUp() {
// 	// 程序员老得太快 ^_^
// 	p.age += 10
// 	return
// }

// func main() {
// 	num := reflect.New(reflect.TypeOf(1))
// 	fmt.Println(num.Kind())
// 	inst := num.Interface().(int)
// 	fmt.Println(num, inst)
// }

// func main() {
// 	var wg sync.WaitGroup
// 	typ := reflect.TypeOf(&wg)
// 	for i := 0; i < typ.NumMethod(); i++ {
// 		method := typ.Method(i)
// 		argv := make([]string, 0, method.Type.NumIn())
// 		returns := make([]string, 0, method.Type.NumOut())

// 		for j := 1; j < method.Type.NumIn(); j++ {
// 			argv = append(argv, method.Type.In(j).Name())
// 		}
// 		for j := 0; j < method.Type.NumOut(); j++ {
// 			returns = append(returns, method.Type.Out(j).Name())
// 		}

// 		log.Printf("func (w *%s) %s(%s)%s",
// 			typ.Elem().Name(),
// 			method.Name,
// 			strings.Join(argv, ","),
// 			strings.Join(returns, ","))
// 	}
// }

// func prints(i int) string {
// 	fmt.Println("i:=", i)
// 	return strconv.Itoa(i)
// }

// func main() {
// 	fv := reflect.ValueOf(prints)
// 	params := make([]reflect.Value, 1)
// 	params[0] = reflect.ValueOf(20)
// 	rs := fv.Call(params)
// 	fmt.Println("result", rs[0].Interface().(string))
// }

// func sum(i, j int) int {
// 	return i + j
// }

// func main() {
// 	fv := reflect.ValueOf(sum)
// 	params := make([]reflect.Value, 2)
// 	params[0] = reflect.ValueOf(10)
// 	params[1] = reflect.ValueOf(20)
// 	rs := fv.Call(params)
// 	fmt.Println("result:", rs[0].Interface().(int))
// }

// type Person struct {
// 	Name string
// 	Age  int
// }

// func (p Person) SayHello(a int) int {
// 	fmt.Println("hello,", p.Name, p.Age)
// 	return 0
// }

// // func main() {
// // 	p := Person{Name: "alice", Age: 30}
// // 	v := reflect.ValueOf(p)
// // 	m := v.MethodByName("SayHello")
// // 	m.Call(nil)
// // }

// func main() {
// 	p := Person{Name: "alice", Age: 30}
// 	v := reflect.ValueOf(p)
// 	// typ := v.Type()
// 	fmt.Println(v.NumMethod(), v.Type().Name())
// 	for i := 0; i < v.NumMethod(); i++ {
// 		met := v.Method(i)
// 		// met.Call(nil)
// 		typ := met.Type()

// 		fmt.Println(typ.Name(), typ.NumIn(), typ.NumOut())
// 		for j := 0; j < typ.NumIn(); j++ {
// 			fmt.Println(typ.In(j))
// 		}

// 		for j := 0; j < typ.NumOut(); j++ {
// 			fmt.Println(typ.Out(j))
// 		}
// 		log.Println("end", met.Type().Name())
// 		res := met.Call([]reflect.Value{reflect.ValueOf(1)})
// 		fmt.Println(res[0].Interface().(int))
// 	}
// }

// func TestFunc(a int) int {
// 	fmt.Println("This is a test function.")
// 	return 0
// }

// func main() {
// 	funcValue := reflect.ValueOf(TestFunc)

// 	// funcName := funcValue.Type().Name()
// 	fmt.Println(funcValue, funcValue.Type(), funcValue.Type().Name())
// }

// type Person struct {
// 	Name string
// 	Age  int
// }

// func (p Person) SayHello(a int) int {
// 	fmt.Printf("Hello, my name is %s and I am %d years old\n", p.Name, p.Age)
// 	fmt.Println(a)
// 	return 0
// }

// func main() {
// 	p := Person{Name: "Alice", Age: 30}

// 	// 使用反射调用结构体方法
// 	v := reflect.ValueOf(p)
// 	m := v.MethodByName("SayHello")
// 	// fmt.Println(m.Type().Method())
// 	// // m.Call(nil)

// 	// fmt.Println(m.Name())
// 	// fmt.Println(m.Addr().Type().Name(), m.Elem().Type().Name())
// 	// v.Method(0)
// 	// t := reflect.TypeOf(p)
// 	// fmt.Println(t.Method(0).Name)
// 	// res := t.Method(0).Func.Call([]reflect.Value{reflect.ValueOf(p), reflect.ValueOf(1)})
// 	// fmt.Println(res)
// }

type Args struct{ Num1, Num2 int }

func main() {
	// log.Println(reflect.TypeOf(1))
	// log.Println(reflect.New(reflect.TypeOf(1)).Elem().Type())

	// var p *int
	// log.Println(reflect.TypeOf(p))
	// log.Println(reflect.New(reflect.TypeOf(p).Elem()).Type())

	// var args Args
	// log.Println(reflect.New(reflect.TypeOf(args)))
	// val := reflect.New(reflect.TypeOf(args)).Elem()
	// log.Println(val, val.Addr().Interface())

	var i int
	val := reflect.New(reflect.TypeOf(i)).Elem()
	modifyVal(val.Addr().Interface())
	log.Println(val)
}

func modifyVal(i interface{}) {
	val := i.(*int)
	*val = 3
	log.Println(*val)
}
