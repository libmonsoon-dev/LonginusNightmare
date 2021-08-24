package main

import (
	"context"

	"github.com/libmonsoon-dev/LonginusNightmare/run"
	"github.com/libmonsoon-dev/LonginusNightmare/server"
	"github.com/libmonsoon-dev/LonginusNightmare/server/wscontroller"
)

type App struct {
	Hub    *wscontroller.Hub
	Server *server.Server
}

func (a App) Run(ctx context.Context) error {
	g := run.Gather{
		run.Errorf("web socket hub", a.Hub),
		run.Errorf("server", a.Server),
	}
	return g.Run(ctx)
}
