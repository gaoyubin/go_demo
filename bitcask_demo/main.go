package main

import (
	"bitcask_demo/bitcask"
	"fmt"
)

func main() {
	bitcask, err := bitcask.OpenBitCask(".")
	if err != nil {
		panic(err)
	}
	// bitcask.Put([]byte("3"), []byte("hello"))
	// bitcask.Put([]byte("4"), []byte("world"))
	val, err := bitcask.Get([]byte("5"))
	fmt.Println(string(val), err)
}
