package main

import (
	"context"

	"github.com/libmonsoon-dev/LonginusNightmare/client/ui"

	"github.com/libmonsoon-dev/LonginusNightmare/client"
	"github.com/libmonsoon-dev/LonginusNightmare/logger/logrus"
)

func main() {
	log := logrus.NewFactory().New("main")

	c, err := client.New()
	if err != nil {
		log.Errorf("inti client: %v", err)
		return
	}

	// TODO: move to settings
	ui.RequestFullscreen()
	defer ui.ExitFullscreen()

	ctx := context.Background()
	if err = c.Run(ctx); err != nil {
		log.Errorf("runtime error: %v", err)
		return
	}
}
