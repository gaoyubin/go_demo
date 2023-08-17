package mycache

import "fmt"

type LinkedList struct {
	key, val int
	next     *LinkedList
	pre      *LinkedList
}

type LruCache struct {
	cur_cap int
	max_cap int
	head    *LinkedList
	tail    *LinkedList
	m       map[int]*LinkedList
}

func NewCache(max_c int) LruCache {
	cache := LruCache{
		cur_cap: 0,
		max_cap: max_c,
		head:    &LinkedList{},
		tail:    &LinkedList{},
		m:       make(map[int]*LinkedList),
	}
	cache.head.next = cache.tail
	cache.tail.pre = cache.head
	return cache
}
func (cache *LruCache) Get(key int) int {
	if node, ok := cache.m[key]; ok {
		cache.removeNode(node)
		cache.moveToHead(node)
		return node.val
	} else {
		return -1
	}
}
func (cache *LruCache) Put(key, val int) {
	if node, ok := cache.m[key]; ok {
		fmt.Println("put old", node)
		node.val = val
		//放到头节点
		cache.removeNode(node)
		cache.moveToHead(node)
	} else {
		node := &LinkedList{
			key: key,
			val: val,
		}
		fmt.Println("put new", node)
		cache.m[key] = node
		//放到头节点
		cache.moveToHead(node)
		cache.cur_cap++
		if cache.cur_cap > cache.max_cap {
			//删除掉尾节点
			remove_node := cache.tail.pre
			fmt.Println("delete olde", remove_node)
			cache.removeNode(remove_node)
			delete(cache.m, remove_node.key)
			cache.cur_cap--
		}
	}
}

func (cache *LruCache) removeNode(node *LinkedList) {
	node.next.pre = node.pre
	node.pre.next = node.next
}

func (cache *LruCache) moveToHead(node *LinkedList) {
	node.pre = cache.head
	node.next = cache.head.next
	cache.head.next.pre = node
	cache.head.next = node
}
