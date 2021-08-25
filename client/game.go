package client

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/text"

	"github.com/libmonsoon-dev/LonginusNightmare/app"
	"github.com/libmonsoon-dev/LonginusNightmare/client/assets/images"
	"github.com/libmonsoon-dev/LonginusNightmare/client/ui/fonts"
)

const (
	screenWidth  = 320
	screenHeight = 240
)

func NewGame() (*Game, error) {
	arcadeN, err := fonts.NewArcadeN()
	if err != nil {
		return nil, fmt.Errorf("init font: %w", err)
	}

	stdImg, _, err := image.Decode(bytes.NewReader(images.Gopher))
	if err != nil {
		return nil, fmt.Errorf("decode gopher image: %w", err)
	}

	gopherImg, err := ebiten.NewImageFromImage(stdImg, ebiten.FilterDefault)
	if err != nil {
		return nil, fmt.Errorf("convert gopher stdimg: %w", err)
	}

	g := &Game{
		ctx:      context.Background(),
		gopher:   gopherImg,
		font:     arcadeN,
		showFPS:  app.Dev,
		player:   NewPlayer(),
		inputCtl: NewInputController(),
	}

	g.initBackground(screenWidth, screenHeight)
	return g, nil
}

type Game struct {
	ctx        context.Context
	gopher     *ebiten.Image
	font       *fonts.Arcade
	background *ebiten.Image
	showFPS    bool
	player     *Player
	inputCtl   *InputController
}

func (g *Game) Run(ctx context.Context) error {
	ebiten.SetWindowTitle(app.Name)
	ebiten.SetWindowSize(screenWidth, screenHeight)
	g.setContext(ctx)

	err := ebiten.RunGame(g)
	if err != nil {
		return fmt.Errorf("run ebiten game: %w", err)
	}

	return nil
}

func (g *Game) Update(_ *ebiten.Image) (err error) {
	err = g.ctx.Err()
	if err != nil {
		return err
	}

	g.inputCtl.Update()
	g.player.Update(g.inputCtl.State())
	return
}

func (g *Game) Draw(screen *ebiten.Image) {
	err := g.background.Clear()
	if err != nil {
		//	TODO: error handling
		fmt.Println(err)
		return
	}
	text.Draw(g.background, "Hello World!", g.font.Title(), 110, screenHeight/2, color.White)
	g.drawPlayer()

	err = screen.DrawImage(g.background, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	if g.showFPS {
		err = ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %0.2f\nTPS: %0.2f", ebiten.CurrentFPS(), ebiten.CurrentTPS()))
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	currentWidth, currentHeight := g.background.Size()
	if outsideWidth != currentWidth || outsideHeight != currentHeight {
		g.initBackground(outsideWidth, outsideHeight)
	}

	// TODO: remove all fmt.Print* calls and migrate to logger
	fmt.Printf("outsideWidth = %d, outsideHeight = %d\n", outsideWidth, outsideHeight)
	return screenWidth, screenHeight
}

func (g *Game) drawPlayer() {
	op := new(ebiten.DrawImageOptions)
	op.GeoM.Scale(0.1, 0.1)
	op.GeoM.Translate(float64(g.player.Position.X), float64(g.player.Position.Y))

	//fmt.Printf("gopher X = %d, Y = %d\n", g.player.Position.X, g.player.Position.Y)
	err := g.background.DrawImage(g.gopher, op)
	if err != nil {
		fmt.Println("draw player", err)
	}
}

func (g *Game) initBackground(width int, height int) {
	background, err := ebiten.NewImage(width, height, ebiten.FilterDefault)
	if err != nil {
		fmt.Println(fmt.Errorf("init background image: %w", err))
		return
	}

	g.background = background
}

func (g *Game) setContext(ctx context.Context) {
	g.ctx = ctx
}
