package main

import (
	_ "net/http/pprof"
)

// "net/http"
// _ "net/http/pprof"
// "strconv"
// "strings"

// func fib(n int) int {
// 	if n <= 1 {
// 		return 1
// 	}
// 	return fib(n-1) + fib(n-2)
// }

// func main() {
// 	// f, _ := os.OpenFile("cpu.profile", os.O_CREATE|os.O_RDWR, 0777)
// 	// defer f.Close()
// 	// pprof.StartCPUProfile(f)
// 	// defer pprof.StopCPUProfile()
// 	// defer profile.Start().Stop()
// 	defer profile.Start().Stop()
// 	n := 10
// 	for i := 1; i < 6; i++ {
// 		log.Println(n, fib(n))
// 		n += 3 * i

// 	}
// }

/*
const Letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func generate(n int) string {
	var buf bytes.Buffer
	for i := 0; i < n; i++ {
		buf.WriteByte(Letters[rand.Intn(len(Letters))])
	}
	return buf.String()
}

func repeatestring(s string, n int) string {
	// var result string
	// for i := 0; i < n; i++ {
	// 	result += s
	// }
	// return result

	// buf := make([]byte, 0, n*len(s))
	// for i := 0; i < n; i++ {
	// 	buf = append(buf, []byte(s)...)
	// }
	// return string(buf)

	buf := &bytes.Buffer{}
	for i := 0; i < n; i++ {
		buf.Write([]byte(s))
	}
	return buf.String()
}
func main() {
	f, err := os.OpenFile("mem.profile", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	for i := 0; i < 100; i++ {
		str := generate(100)
		repeatestring(str, 1000)
	}
	// defer pprof.StopCPUProfile()
	pprof.Lookup("heap").WriteTo(f, 0)
	// pprof.StartCPUProfile(f)
	// time.Sleep(3 * time.Second)
	// bytes := make([]byte, 11)
	// // log.Println(len(bytes))
	// // size := copy(bytes, "hello")
	// // copy(bytes[size:], " world")
	// bytes = append(bytes, []byte("hello")...)
	// fmt.Println(string(bytes), len(bytes))
}
*/

// func fib(n int) int {
// 	if n <= 1 {
// 		return 1
// 	}
// 	return fib(n-1) + fib(n-2)
// }

// func fibHandler(w http.ResponseWriter, r *http.Request) {
// 	n, err := strconv.Atoi(r.URL.Path[len("/fib/"):])
// 	if err != nil {
// 		// http.ResponseWriter
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// 	var result int
// 	for i := 0; i < 100; i++ {
// 		result = fib(n)
// 	}
// 	// binary.Write()
// 	w.Write([]byte(strconv.Itoa(result)))
// }
// func repeat(s string, n int) string {
// 	var result string
// 	for i := 0; i < n; i++ {
// 		result += s
// 	}

// 	return result
// }
// func repeatHandler(w http.ResponseWriter, r *http.Request) {
// 	suffix := r.URL.Path[len("/repeat/"):]
// 	parts := strings.SplitN(suffix, "/", 2)
// 	if len(parts) != 2 {
// 		http.Error(w, "invalid param", http.StatusBadRequest)
// 		return
// 	}
// 	s := parts[0]
// 	n, err := strconv.Atoi(parts[1])
// 	if err != nil {
// 		http.Error(w, "invalid param", http.StatusBadRequest)
// 		return
// 	}
// 	var result string
// 	for i := 0; i < 1000; i++ {
// 		result = repeat(s, n)
// 	}
// 	w.Write([]byte(result))
// }

// func main() {
// 	// http.HandleFunc();
// 	// u, err := url.Parse("https://example.org:8000/path")
// 	// if err != nil {
// 	// 	panic(err)
// 	// }
// 	// fmt.Println(u.Path, u.Host)
// 	// name := "John"
// 	// age := 30
// 	// height := 175.5

// 	// result := fmt.Sprintf("Name: %v, Age: %v, Height: %v", name, age, height)
// 	// fmt.Println(result)

// 	go func() {
// 		//默认的mux
// 		log.Fatal(http.ListenAndServe(":9090", nil))
// 	}()

// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/fib/", fibHandler)
// 	mux.HandleFunc("/repeat/", repeatHandler)
// 	s := &http.Server{
// 		Addr:    ":8080",
// 		Handler: mux,
// 	}

// 	//新的mux
// 	if err := s.ListenAndServe(); err != nil {
// 		log.Fatal(err)
// 	}
// 	// l, err := net.Listen("tcp", ":8080")
// 	// if err != nil {
// 	// 	panic(err)
// 	// }
// 	// http.Serve(l)
// 	// NewProfile
// 	// log.Fatal(http.ListenAndServe(":8080", nil))
// }
