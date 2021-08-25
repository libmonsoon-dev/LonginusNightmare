package client

import (
	"github.com/hajimehoshi/ebiten"

	v1 "github.com/libmonsoon-dev/LonginusNightmare/models/v1"
)

func NewInputController() *InputController {
	ctl := &InputController{
		state: v1.NewUserInput(),
	}
	return ctl
}

type InputController struct {
	state *v1.UserInput
}

func (c *InputController) Update() {
	c.updateDir()
}

func (c *InputController) State() *v1.UserInput {
	return c.state
}

func (c *InputController) updateDir() {
	c.state.Direction = c.getDir()
}

func (c *InputController) getDir() v1.Direction {
	switch {
	case ebiten.IsKeyPressed(ebiten.KeyUp):
		return v1.DirectionUp

	case ebiten.IsKeyPressed(ebiten.KeyLeft):
		return v1.DirectionLeft

	case ebiten.IsKeyPressed(ebiten.KeyRight):
		return v1.DirectionRight

	case ebiten.IsKeyPressed(ebiten.KeyDown):
		return v1.DirectionDown
	}

	return v1.DirectionNone
}
