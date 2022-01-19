package lru

import "container/list"


type Cache struct {
	maxBytes int64
	nbytes   int64
	ll       *list.List
	cache    map[string]*list.List
	onEvicted func(key string, value Value)
}


type entry struct {
	key   string
	value Value
}


type Value interface {
	Len() int
}


func New(maxBytes int64, onEvicted func(string, Value)) *Cache {
	
}