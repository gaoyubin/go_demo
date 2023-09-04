package main

import (
	"fmt"
	"log"
	"time"

	"github.com/go-zookeeper/zk"
	// "github.com/go-zookeeper/zk@v1.0.3"
)

var zkconn *zk.Conn

func init() {
	var err error
	zkconn, _, err = zk.Connect([]string{"127.0.0.1"}, 2*time.Second)
	// zk.CreateContainerRequest
	// fmt.Append()
	if err != nil {
		panic(err)
	}
	// defer zkconn.Close()
	// log.Println(zkconn.)
}

func main() {
	// res, stat, err := zkconn.Get("/app1/p_1")
	// if err != nil {
	// 	panic(err)
	// }
	// log.Println(string(res), stat)
	// log.Println("succ")

	// zkconn.Create("/app2/lock", nil, )

	// var wg sync.WaitGroup
	// for i := 0; i < 50; i++ {
	// 	wg.Add(1)
	// 	go func(n int) {
	// 		defer wg.Done()
	// 		lock := zk.NewLock(zkconn, "/app3/lock", zk.WorldACL(zk.PermAll))
	// 		err := lock.LockWithData([]byte("it is a lock"))
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 		log.Println("第", n, "个goroutine获取到了锁")
	// 		time.Sleep(time.Second)
	// 		lock.Unlock()
	// 	}(i)
	// }
	// wg.Wait()

	_, _, eventChannel, err := zkconn.ExistsW("/app4")
	if err != nil {
		panic(err)
	}

	go func() {
		e := <-eventChannel
		log.Println("watch", e)
	}()

	var ss string
	fmt.Scanln(&ss)

}
