package main

import (
	"fmt"
	"lru_cache_demo/mycache"
)

func main() {
	fmt.Println("hello")
	cache := mycache.NewCache(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Put(3, 3)
	fmt.Println(cache.Get(1))
	x := mycache.TestInterface(&cache, 1)
	fmt.Println(x)
}
