package main

//
// type Node struct {
// 	Pre  *Node
// 	Next *Node
// }
// type LruCache struct {
// 	head    *Node
// 	tail    *Node
// 	m       map[int]*Node
// 	cap     int
// 	curSize int
// }
//
// func NewLruCache(cap int) *LruCache {
// 	h := &Node{}
// 	t := &Node{}
// 	h.Pre = t
// 	h.Next = t
// 	t.Pre = h
// 	t.Next = h
//
// 	return &LruCache{
// 		m:       make(map[int]*Node),
// 		cap:     cap,
// 		curSize: 0,
// 		head:    h,
// 		tail:    t,
// 	}
//
// }

//
// func (cache *LruCache) put(key int, val int) {
// 	if _, ok := cache.m[key]; ok {
// 		n := cache.m[key]
// 		// 把node挪到head
// 	} else {
//
// 	}
// }

// func main() {
//
// 	count := 0
// 	fmt.Scan(&count)
// 	fmt.Println(count)
// 	l := []int{}
// 	for i := 0; i < count; i++ {
// 		tmp := 0
// 		fmt.Scan(&tmp)
// 		l = append(l, tmp)
// 	}
// 	fmt.Println(l)
// }
