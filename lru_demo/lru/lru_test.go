package lru

import (
	"fmt"
	"testing"
)

type Int int

func (d Int) Len() int {
	return 1
}

func TestGet(t *testing.T) {
	lru := NewCache(4)
	lru.Put("1", Int(1))
	lru.Put("2", Int(2))
	// if v, ok := lru.Get("2");!ok{
	// 	t.Fatal("cache hit fail", "2")
	// }
	v, ok := lru.Get("2")
	fmt.Println("hit cache", v, ok)
	lru.Put("3", Int(3))

	v, ok = lru.Get("1")
	fmt.Println("hit cache", v, ok)
}
