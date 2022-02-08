package geecache


// ByteView 抽象的只读数据结构，用于表示缓存值
type ByteView struct {
	b []byte
}


func (v ByteView) ByteSlice() []byte {
	return cloneBytes(v.b)
}


func (v ByteView) String() string {
	return string(v.b)
}


func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(b, c)
	return c
}
