package pool

import (
	"bytes"
	"sync"
)

var joinStringsPool = sync.Pool{
	New: func() any {
		return &bytes.Buffer{}
	},
}

func GetBuffer() *bytes.Buffer {
	buf := joinStringsPool.Get().(*bytes.Buffer)
	buf.Reset()

	return buf
}

func PutBuffer(b *bytes.Buffer) {
	joinStringsPool.Put(b)
}
