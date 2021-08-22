package server

import (
	"net/http"

	"github.com/valyala/fasthttp"
)

func httpErr(ctx *fasthttp.RequestCtx, status int) {
	ctx.Error(http.StatusText(status), status)
}
