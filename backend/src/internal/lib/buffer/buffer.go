package buffer

import "sync"

const (
	kb = 1 << 10

	maxBufferSize = 16 << 10
)

type Buffer []byte

var (
	bufPool = sync.Pool{
		New: func() any {
			b := make(Buffer, 0, kb)
			return &b
		},
	}
)

func New() *Buffer {
	return bufPool.Get().(*Buffer)
}

func (b *Buffer) Free() {
	if cap(*b) <= maxBufferSize {
		*b = (*b)[:0]
		bufPool.Put(b)
	}
}

func (b *Buffer) Len() int {
	return len(*b)
}

func (b *Buffer) String() string {
	return string(*b)
}

func (b *Buffer) WriteString(s string) {
	*b = append(*b, s...)
}
