package wscontroller

import (
	"github.com/fasthttp/websocket"

	"github.com/libmonsoon-dev/LonginusNightmare/logger"
)

func New(hub *Hub, loggerFactory logger.Factory) func(conn *websocket.Conn) {
	c := &controller{
		loggerFactory: loggerFactory,
		logger:        loggerFactory.New("WebSocket hub"),
		hub:           hub,
	}
	return c.Handle
}

type controller struct {
	loggerFactory logger.Factory
	logger        logger.Logger
	hub           *Hub
}

func (c *controller) Handle(conn *websocket.Conn) {
	session := c.newSession(conn)
	session.hub.registerCh <- session

	go session.writePump()
	session.readPump()
}
