package lru

import (
	"container/list"
	"fmt"
)

type Cache struct {
	ll      *list.List
	m       map[string]*list.Element
	max_cap int
	cur_cap int
}
type Value interface {
	Len() int
}
type entry struct {
	key   string
	value Value
}

func NewCache(max_cap_ int) *Cache {
	return &Cache{
		ll:      list.New(),
		m:       make(map[string]*list.Element),
		max_cap: max_cap_,
		cur_cap: 0,
	}
}
func (c *Cache) Get(key string) (value Value, ok bool) {
	if elem, ok := c.m[key]; ok {
		c.ll.MoveToFront(elem)
		kv := elem.Value.(*entry)
		return kv.value, true
	}
	return
}

func (c *Cache) Put(key string, value Value) bool {
	if elem, ok := c.m[key]; ok {
		c.ll.MoveToFront(elem)
		kv := elem.Value.(*entry)
		c.cur_cap += value.Len() - kv.value.Len()
		kv.value = value
	} else {
		elem := c.ll.PushFront(&entry{key, value})
		c.m[key] = elem
		c.cur_cap += len(key) + value.Len()
	}
	c.MoveOldest()
	return true
}
func (c *Cache) MoveOldest() {
	for c.cur_cap > c.max_cap {
		val := c.ll.Remove(c.ll.Back())
		kv := val.(*entry)
		delete(c.m, kv.key)
		fmt.Println("MoveOldest", kv, c.cur_cap)
		c.cur_cap -= (len(kv.key) + kv.value.Len())
	}
}
