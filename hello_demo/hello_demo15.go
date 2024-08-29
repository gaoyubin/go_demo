package main

import "fmt"

const loadFactor = 1 // 负载因子(决定扩容因数)

type Entry struct {
	key  string
	val  interface{}
	next *Entry
}
type HashMap struct {
	size   int
	bucket []Entry
}

func NewHashMap() *HashMap {
	hm := &HashMap{
		size:   0,
		bucket: make([]Entry, 4),
	}
	return hm
}
func (hm *HashMap) Put(key string, val interface{}) {
	e := Entry{
		key: key,
		val: val,
	}
	hm.insert(e)
	fmt.Println(hm)
	if float32(hm.size)/float32(len(hm.bucket)) >= loadFactor {
		fmt.Println("hit extend")
		newHm := HashMap{
			size:   0,
			bucket: make([]Entry, len(hm.bucket)*2),
		}
		for _, e := range hm.bucket {
			if e.key == "" {
				continue
			} else {
				for e.next != nil {
					newHm.insert(e)
					e = *e.next
				}
				newHm.insert(e)
			}
		}
		*hm = newHm
	}

}

func (hm *HashMap) Get(key string) interface{} {
	index := hm.hashcode(key, len(hm.bucket))
	e := &hm.bucket[index]
	if e.key == "" {
		return nil
	} else {
		for e != nil {
			if e.key == key {
				return e.val
			}
		}
		return nil
	}
}
func (hm *HashMap) insert(entry Entry) {
	index := hm.hashcode(entry.key, len(hm.bucket))
	e := &hm.bucket[index]
	if e.key == "" {
		*e = entry
		hm.size++
		return
	} else {
		for e != nil {
			if e.key == entry.key {
				*e = entry
				return
			}
			e = e.next
		}

		e.next = &entry
		hm.size++

	}

}

func (hm *HashMap) hashcode(key string, len int) int {
	sum := 0
	for i, _ := range key {
		sum += int(key[i])
	}
	return sum % len
}
func main() {
	hm := NewHashMap()
	fmt.Println(hm)
	hm.Put("1", "hello")
	hm.Put("2", "world")
	hm.Put("3", 12)
	hm.Put("4", 45)
	fmt.Println(hm.Get("2"))
	fmt.Println(hm.Get("3"))
	hm.Put("5", "echo")
	fmt.Println(hm.Get("3"))
}
