package main

import (
	"log"

	"github.com/valyala/fasthttp"

	"github.com/libmonsoon-dev/LonginusNightmare/server"
	"github.com/libmonsoon-dev/LonginusNightmare/static"
)

func main() {
	addr := "0.0.0.0:1337"

	handler := server.NewStaticHandler(static.FS)
	log.Println("starting server on", addr)
	if err := fasthttp.ListenAndServe(addr, handler); err != nil {
		log.Fatal(err)
	}
}
