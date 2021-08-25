package client

import (
	"image"

	v1 "github.com/libmonsoon-dev/LonginusNightmare/models/v1"
)

type Player struct {
	Position image.Point
	Speed    int
}

func NewPlayer() *Player {
	p := &Player{
		Speed:    2,
		Position: image.Point{X: 130, Y: 125},
	}

	return p
}

func (p *Player) Update(state *v1.UserInput) {
	p.updatePosition(state.Direction)
}

func (p *Player) updatePosition(dir v1.Direction) {
	if dir == v1.DirectionNone {
		return
	}

	switch dir {
	case v1.DirectionNone:
		return

	case v1.DirectionUp:
		p.Position.Y -= p.Speed

	case v1.DirectionLeft:
		p.Position.X -= p.Speed

	case v1.DirectionRight:
		p.Position.X += p.Speed

	case v1.DirectionDown:
		p.Position.Y += p.Speed
	}
}
