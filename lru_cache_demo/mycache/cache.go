package mycache

type Cache interface {
	Get(key int) int
	Put(key, val int)
}

func TestInterface(cache Cache, key int) int {
	return cache.Get(key)
	// return 1
}
