package singleflight

import "sync"


// call 代表正在进行中或已经结束的请求
type call struct {
	wg  sync.WaitGroup  // 避免重入锁；不需要消息传递的并发协程
	val interface{}
	err error
}


// Group 管理不同 key 的请求 call
type Group struct {
	mu sync.Mutex  // 保护 m 不被并发读写的锁
	m  map[string]*call
}


// Do 针对相同的key，函数fn只调用一次
func (g *Group) Do(key string, fn func() (interface{}, error)) (interface{}, error) {
	g.mu.Lock()
	if g.m == nil {
		g.m = make(map[string]*call)
	}

 	// 如果有请求正在进行，则阻塞等待，直到锁释放
	if c, ok := g.m[key]; ok {
		g.mu.Unlock()
		c.wg.Wait() 
		return c.val, c.err
	}

	c := new(call)
	c.wg.Add(1)  // 锁 +1
	g.m[key] = c
	g.mu.Unlock()

	c.val, c.err = fn()
	c.wg.Done()  // 锁 -1

	g.mu.Lock()
	delete(g.m, key)  // 更新 g.m，删除key
	g.mu.Unlock()

	return c.val, c.err
}
