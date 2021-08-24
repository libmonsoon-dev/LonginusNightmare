package httpcontroller

import (
	"net/http"

	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

func (c *controller) index(ctx *fasthttp.RequestCtx) {
	ctx.Success("text/html; charset=UTF-8", c.indexPage)
}

func (c *controller) static() fasthttp.RequestHandler {
	stdHandler := http.FileServer(http.FS(c.staticFs))
	stdHandler = http.StripPrefix("/static/", stdHandler)

	return fasthttpadaptor.NewFastHTTPHandler(stdHandler)
}

func (c *controller) wsHandle(ctx *fasthttp.RequestCtx) {
	err := c.upgrader.Upgrade(ctx, c.wsHandler)
	if err != nil {
		c.logger.Errorf("upgrade to WebSocket: %v", err)
	}
}
