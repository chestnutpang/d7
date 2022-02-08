package geecache

// 支持并发读写，数据缓存

type ByteView struct {
	b []byte
}

func (v ByteView) ByteSlice() []byte {
	return cloneByte(v.b)
}


func (v ByteView) String() string {
	return string(v.b)
}


func cloneByte(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
