package server

import (
	"context"
	"fmt"

	"github.com/valyala/fasthttp"

	"github.com/libmonsoon-dev/LonginusNightmare/app"
)

func NewServer(handler fasthttp.RequestHandler, config *Config) *Server {
	s := &Server{
		config: config,
		server: &fasthttp.Server{
			Name:    app.Name,
			Handler: handler,
		},
	}
	return s
}

type Server struct {
	config *Config
	server *fasthttp.Server
}

func (s *Server) Run(ctx context.Context) error {
	serveErr := make(chan error, 1)
	go func() { serveErr <- s.server.ListenAndServe(s.config.Addr) }()

	select {
	case <-ctx.Done():
		s.server.Shutdown()
		return ctx.Err()
	case err := <-serveErr:
		return fmt.Errorf("listen and serve: %w", err)
	}
}
