package httpcontroller

import (
	"io/fs"

	"github.com/fasthttp/router"
	"github.com/fasthttp/websocket"
	"github.com/valyala/fasthttp"

	"github.com/libmonsoon-dev/LonginusNightmare/logger"
)

func New(logFactory logger.Factory, indexPage []byte, staticFs fs.FS, upgrader *websocket.FastHTTPUpgrader,
	wsHandler func(conn *websocket.Conn)) fasthttp.RequestHandler {
	c := &controller{
		logger:    logFactory.New("http controller"),
		router:    router.New(),
		indexPage: indexPage,
		staticFs:  staticFs,
		upgrader:  upgrader,
		wsHandler: wsHandler,
	}
	c.Init()

	return c.router.Handler
}

type controller struct {
	logger    logger.Logger
	router    *router.Router
	indexPage []byte
	staticFs  fs.FS
	upgrader  *websocket.FastHTTPUpgrader
	wsHandler func(conn *websocket.Conn)
}

func (c *controller) Init() {
	c.router.GET("/", c.index)
	c.router.GET("/static/{filepath:*}", c.static())
	c.router.GET("/ws", c.wsHandle)
}
