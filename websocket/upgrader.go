package websocket

import (
	"github.com/fasthttp/websocket"
	"github.com/valyala/fasthttp"
)

func NewUpgrader() *websocket.FastHTTPUpgrader {
	const megabyte = 1 << 20
	poolWriteBuffer := newBufferPool(4096, megabyte)

	return &websocket.FastHTTPUpgrader{
		WriteBufferPool:   poolWriteBuffer,
		ReadBufferSize:    4096,
		WriteBufferSize:   4096,
		EnableCompression: true,
		CheckOrigin: func(ctx *fasthttp.RequestCtx) bool {
			return true
		},
	}
}
