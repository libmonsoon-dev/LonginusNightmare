package client

import (
	"context"
	"fmt"

	"github.com/libmonsoon-dev/LonginusNightmare/run"
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

func (c *Client) Run(ctx context.Context) error {
	return run.Errorf("game", c.game).Run(ctx)
}
