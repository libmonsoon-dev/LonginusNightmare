package client

import (
	"fmt"

	"github.com/hajimehoshi/ebiten"

	"github.com/libmonsoon-dev/LonginusNightmare/app"
)

type Client struct {
	game *Game
}

func New() (*Client, error) {
	game, err := NewGame()
	if err != nil {
		return nil, fmt.Errorf("init game: %w", err)
	}

	c := &Client{
		game: game,
	}
	return c, nil
}

func (c *Client) Run() error {
	ebiten.SetWindowTitle(app.Name)
	ebiten.SetWindowSize(screenWidth, screenHeight)
	return ebiten.RunGame(c.game)
}
