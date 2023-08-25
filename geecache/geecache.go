package geecache

import (
	"fmt"
	"log"
	"sync"
)

type Group struct {
	name      string
	getter    Getter
	mainCache cache
}
type Getter interface {
	Get(key string) ([]byte, error)
}
type GetterFunc func(key string) ([]byte, error)

func (f GetterFunc) Get(key string) ([]byte, error) {
	return f(key)
}

var (
	mu     sync.Mutex
	groups = make(map[string]*Group)
)

func NewGroup(name string, cacheBytes int, getter Getter) *Group {
	mu.Lock()
	defer mu.Unlock()
	if getter == nil {
		panic("no getter")
	}
	g := &Group{
		name:   name,
		getter: getter,
		mainCache: cache{
			cacheBytes: cacheBytes},
	}
	groups[name] = g
	return g
}
func GetGroup(name string) *Group {
	mu.Lock()
	defer mu.Unlock()
	g := groups[name]
	return g
}

func (g *Group) Get(key string) (ByteView, error) {
	if key == "" {
		return ByteView{}, fmt.Errorf("key is required")
	}
	if v, ok := g.mainCache.get(key); ok {
		log.Println("geecache hit", key, v)
		return v, nil
	}
	return g.load(key)
}
func (g *Group) load(key string) (value ByteView, err error) {
	// bytes, err := g.getter.Get(key)
	// if err != nil {
	// 	return ByteView{}, err
	// }
	// value = ByteView{b: bytes}
	// g.mainCache.add(key, value)
	// return value, nil
	//请求与远端服务器
	if g.peers != nil {
		if peer, ok := g.peers.PickPeer(key); ok {
			bytes, err := peer.Get(g.name, key)
			if err != nil {
				return ByteView{}, err
			}
			return ByteView{b: bytes}, nil
		}
	}
	//找不到，请求本地
	return g.getLocally(key)
}

func (g *Group) getLocally(key string) (ByteView, error) {
	bytes, err := g.getter.Get(key)
	if err != nil {
		return ByteView{}, err
	}
	value := ByteView{b: bytes}
	//加载到cache
	g.mainCache.add(key, value)
	return value, nil
}

// func (g *Group) getFromPeer(peer PeerGetter, key string)(byte)
