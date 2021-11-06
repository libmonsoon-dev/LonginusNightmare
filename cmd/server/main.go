package main

import (
	"context"
	"flag"
	"os/signal"
	"syscall"

	"github.com/libmonsoon-dev/LonginusNightmare/logger/logrus"
	"github.com/libmonsoon-dev/LonginusNightmare/server"
	"github.com/libmonsoon-dev/LonginusNightmare/server/httpcontroller"
	"github.com/libmonsoon-dev/LonginusNightmare/server/wscontroller"
	"github.com/libmonsoon-dev/LonginusNightmare/static"
	"github.com/libmonsoon-dev/LonginusNightmare/websocket"
)

func main() {
	addr := flag.String("addr", "0.0.0.0:1337", "Address to serve")

	flag.Parse()

	serverConfig := &server.Config{
		Addr: *addr,
	}

	var app App
	logFactory := logrus.NewFactory()
	app.Hub = wscontroller.NewHub()
	wsCtl := wscontroller.New(app.Hub, logFactory)
	upgrader := websocket.NewUpgrader()
	httpCtl := httpcontroller.New(logFactory, static.Index, static.Static, upgrader, wsCtl)
	app.Server = server.NewServer(httpCtl, serverConfig)

	ctx, stopNotify := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stopNotify()

	err := app.Run(ctx)
	if err != nil {
		logFactory.New("main").Error(err)
	}
}
