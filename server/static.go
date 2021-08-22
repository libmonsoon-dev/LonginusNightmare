package server

import (
	"net/http"

	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

func NewStaticHandler(fs FS) fasthttp.RequestHandler {
	// TODO: migrate to fasthttp.FS
	return fasthttpadaptor.NewFastHTTPHandler(http.FileServer(http.FS(fs)))
}
