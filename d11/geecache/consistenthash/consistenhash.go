package consistenthash


import (
	"hash/crc32"
	"sort"
	"strconv"
)


type Hash func(data []byte) uint32


type Map struct {
	hash     Hash
	replicas int  // 虚拟节点倍数
	keys     []int  // 哈希环
	hashMap  map[int]string  // 虚拟节点与真实节点映射表 key：虚拟节点hash value：真实节点名称
}


func New(replicas int, fn Hash) *Map {
	m := &Map{
		replicas: replicas,
		hash:     fn,
		hashMap:  make(map[int]string),
	}
	if m.hash == nil {
		m.hash = crc32.ChecksumIEEE
	}
	return m
}


// Add 添加真实节点方法
func (m *Map) Add(keys ...string) {
	for _, key := range keys {
		for i := 0; i < m.replicas; i++ {
			hash := int(m.hash([]byte(strconv.Itoa(i) + key)))
			m.keys = append(m.keys, hash)
			m.hashMap[hash] = key
		}
	}
	sort.Ints(m.keys)
}


// Get 获取最接近的元素
func (m *Map) Get(key string) string {
	if len(m.keys) == 0 {
		return ""
	}

	hash := int(m.hash([]byte(key)))

	idx := sort.Search(len(m.keys), func (i int) bool {
		return m.keys[i] >= hash
	})

	return m.hashMap[m.keys[idx % len(m.keys)]]
}