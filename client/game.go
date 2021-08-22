package client

import (
	"fmt"
	"image/color"

	"github.com/libmonsoon-dev/LonginusNightmare/client/ui/font"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/text"
)

const (
	screenWidth  = 320
	screenHeight = 240
)

func NewGame() (*Game, error) {
	f, err := font.NewArcade()
	if err != nil {
		return nil, fmt.Errorf("init font: %w", err)
	}

	background, _ := ebiten.NewImage(screenWidth, screenHeight, ebiten.FilterDefault)
	text.Draw(background, "Hello World!", f.Title(), 110, screenHeight/2, color.White)

	g := &Game{
		font:       f,
		background: background,
		showFPS:    true,
	}

	return g, nil
}

type Game struct {
	font       *font.Arcade
	background *ebiten.Image
	showFPS    bool
}

func (g *Game) Update(_ *ebiten.Image) error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.background, nil)

	if g.showFPS {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %0.2f\nTPS: %0.2f", ebiten.CurrentFPS(), ebiten.CurrentTPS()))
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
