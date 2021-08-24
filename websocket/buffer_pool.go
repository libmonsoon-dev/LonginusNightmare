package websocket

import (
	"sync"
	"unsafe"

	"github.com/fasthttp/websocket"
)

func newBufferPool(initSize, maxSize int) websocket.BufferPool {
	return &bufferPool{maxSize, sync.Pool{
		New: func() interface{} {
			return make([]byte, 0, initSize)
		},
	}}
}

type bufferPool struct {
	maxSize int
	sync.Pool
}

func (b *bufferPool) Put(iface interface{}) {
	type Iface struct {
		Type unsafe.Pointer
		Data *[]byte
	}

	buf := *(*Iface)(unsafe.Pointer(&iface)).Data
	if cap(buf) < b.maxSize {
		b.Pool.Put(iface)
	}
}
