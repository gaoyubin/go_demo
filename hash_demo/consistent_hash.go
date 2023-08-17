package consistent_hash

import (
	"hash/crc32"
	"sort"
	"strconv"
)

type Hash func(data []byte) uint32
type ConsistentHash struct {
	hash     Hash
	replicas int
	keys     []int
	hashMap  map[int]string
}

func New(replicas int, fn Hash) *ConsistentHash {
	consistent_hash := &ConsistentHash{
		replicas: replicas,
		hash:     fn,
		hashMap:  make(map[int]string),
	}
	if consistent_hash.hash != nil {
		consistent_hash.hash = crc32.ChecksumIEEE
	}
	return consistent_hash
}

func (consistent_has *ConsistentHash) Add(keys ...string) {
	for _, key := range keys {
		for i := 0; i < consistent_has.replicas; i++ {
			hash := int(consistent_has.hash([]byte(strconv.Itoa(i) + key)))
			consistent_has.keys = append(consistent_has.keys, hash)
			consistent_has.hashMap[hash] = key
		}
	}
	sort.Ints(consistent_has.keys)
}
