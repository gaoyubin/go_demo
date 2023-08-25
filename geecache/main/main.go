package main

import (
	"flag"
	"fmt"
	"geecache"
	"log"
	"net/http"
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
// var counter = 0
// var lock sync.Mutex

// func Count() {
// 	lock.Lock()
// 	defer lock.Unlock()
// 	counter++

// }
// func main() {
// 	for i := 0; i < 1000; i++ {
// 		go Count()
// 	}
// 	time.Sleep(1500 * time.Microsecond)
// 	fmt.Println(counter)
// }

var db = map[string]string{
	"Tom":  "630",
	"Jack": "589",
	"Sam":  "567",
}

// func main() {
// 	groupName := "scores"
// 	// loadCounts := make(map[string]int, len(db))
// 	gee := geecache.NewGroup(groupName, 2<<10, geecache.GetterFunc(
// 		func(key string) ([]byte, error) {
// 			log.Println("show db", key)
// 			if v, ok := db[key]; ok {
// 				// loadCounts[key] = 0
// 				log.Println("hit db", key, v)
// 				return []byte(v), nil
// 			}
// 			return nil, fmt.Errorf("not find db", key)
// 			// loadCounts[key]++
// 		}))

// 	for k, v := range db {
// 		if view, err := gee.Get(k); err != nil || view.String() != v {
// 			// t.Fatal("failed to get value of Tom")
// 			log.Println("get value fail", k)
// 		} else {
// 			log.Println("get value succ", k, view)
// 		}
// 		// log.Println(view)
// 		// if _, err := gee.Get(k); err != nil || loadCounts[k] > 1 {
// 		// 	// t.Fatalf("cache %s miss", k)
// 		// }
// 	}
// 	// if group := GetGroup(groupName); group == nil || group.name != groupName {
// 	// 	t.Fatalf("group %s not exist", groupName)
// 	// }

// 	// if group := GetGroup(groupName + "111"); group != nil {
// 	// 	t.Fatalf("expect nil, but %s got", group.name)
// 	// }
// }

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	var port int
	var api bool
	flag.IntVar(&port, "port", 8001, "api port")
	flag.BoolVar(&api, "api", false, "start a api server?")
	flag.Parse()

	gee := geecache.NewGroup("scores", 2<<10, geecache.GetterFunc(
		func(key string) ([]byte, error) {
			log.Println("show db search key", key)
			if v, ok := db[key]; ok {
				return []byte(v), nil
			}
			return nil, fmt.Errorf("%s not exist", key)
		}))

	// addr := ":9999"
	// peers := geecache.NewHTTPPool(addr)
	// log.Println("geecache is running at", addr)
	// log.Fatal(http.ListenAndServe(addr, peers))

	addrMap := map[int]string{
		8001: "http://localhost:8001",
		8002: "http://localhost:8002",
		8003: "http://localhost:8003",
	}

	var addrs []string
	for _, v := range addrMap {
		addrs = append(addrs, v)
	}

	apiAddr := "http://localhost:9999"
	if api {
		go startAPIServer(apiAddr, gee)
	}
	startCacheServer(addrMap[port], addrs, gee)
}

func startCacheServer(addr string, addrs []string, gee *geecache.Group) {
	cachesvr := geecache.NewHTTPPool(addr)
	cachesvr.Set(addrs...)
	log.Println("geecache is running at", addr)
	log.Fatal(http.ListenAndServe(addr[7:], cachesvr))

}

func startAPIServer(apiAddr string, gee *geecache.Group) {
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		view, err := gee.Get(key)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Write(view.ByteSlice())
	})
	log.Println("server is running at", apiAddr)
	log.Fatal(http.ListenAndServe(apiAddr[7:], nil))
}
