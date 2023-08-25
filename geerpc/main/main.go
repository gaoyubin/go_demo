package main

import (
	"geerpc"
	"geerpc/registry"
	"geerpc/xclient"
	"log"
	"net"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type Foo int

type Args struct{ Num1, Num2 int }

func (f Foo) Sum(args Args, reply *int) error {
	*reply = args.Num1 + args.Num2
	return nil
}

func startServer() {
	l, err := net.Listen("tcp", ":")
	if err != nil {
		panic(err)
	}
	var foo Foo
	geerpc.Register(&foo)
	// addr <- l.Addr().String()
	log.Println(l.Addr().String())
	registry.HeartBeat(registry_url, l.Addr().String(), 0)
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println("accept err", err)
			continue
		}
		go doServer(conn)
	}

}
func doServer(conn net.Conn) {
	geerpc.DefaultServer.ServeConn(conn)
}

const (
	defaultPath  = "/geerpc/registry"
	defaultPort  = 9999
	registry_url = "http://127.0.0.1:9999/geerpc/registry"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	go func() {
		l, _ := net.Listen("tcp", ":"+strconv.Itoa(defaultPort))
		http.Handle(defaultPath, registry.DefaultGeeRegister)
		http.Serve(l, nil)
	}()
	// addr := make(chan string)

	go startServer()
	go startServer()
	// addr_str := <-addr
	// conn, err := net.Dial("tcp", addr_str)
	// if err != nil {
	// 	panic(err)
	// }
	// log.Println("tcp addr", addr_str)
	// _ = json.NewEncoder(conn).Encode(DefaultOption)
	// cc := codec.NewGobCodec(conn)
	// for i := 0; i < 5; i++ {
	// 	h := &codec.Header{
	// 		ServiceMethod: "Foo.Sum",
	// 		Seq:           uint64(i),
	// 	}
	// 	_ = cc.Write(h, fmt.Sprintf("geerpc req %d", h.Seq))
	// 	_ = cc.ReadHeader(h)
	// 	var reply string
	// 	_ = cc.ReadBody(&reply)
	// 	log.Println("reply:", reply)
	// }
	// // fmt.Scanf("")
	// time.Sleep(10 * time.Second)
	// client, err := geerpc.NewClient(conn)
	time.Sleep(3 * time.Second)
	xclient, err := xclient.NewXClient(xclient.RandomSelect, registry_url)
	if err != nil {
		panic(err)
	}
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			// args := fmt.Sprintf("geerpc req %d", i)
			// var reply string
			// if err := client.Call("Foo.Sum", args, &reply); err != nil {
			// 	log.Fatal("call foo.sum error:", err)
			// }
			args := Args{Num1: i, Num2: i * i}
			var reply int
			log.Println("req:", reply, args)
			if err := xclient.Call("Foo.Sum", args, &reply); err != nil {
				log.Fatal("call foo.sum err", err)
			}
			log.Println("reply:", reply, args)
		}(i)
	}
	wg.Wait()
}
