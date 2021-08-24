package main

import (
	"context"

	"github.com/libmonsoon-dev/LonginusNightmare/logger/logrus"

	"github.com/libmonsoon-dev/LonginusNightmare/server"
	"github.com/libmonsoon-dev/LonginusNightmare/server/httpcontroller"
	"github.com/libmonsoon-dev/LonginusNightmare/server/wscontroller"
	"github.com/libmonsoon-dev/LonginusNightmare/static"
	"github.com/libmonsoon-dev/LonginusNightmare/websocket"
)

func main() {
	serverConfig := &server.Config{
		Addr: "0.0.0.0:1337",
	}

	var app App
	logFactory := logrus.NewFactory()
	app.Hub = wscontroller.NewHub()
	wsCtl := wscontroller.New(app.Hub, logFactory)
	upgrader := websocket.NewUpgrader()
	httpCtl := httpcontroller.New(logFactory, static.Index, static.Static, upgrader, wsCtl)
	app.Server = server.NewServer(httpCtl, serverConfig)

	ctx := context.TODO()
	err := app.Run(ctx)
	if err != nil {
		logFactory.New("main").Error(err)
	}
}
