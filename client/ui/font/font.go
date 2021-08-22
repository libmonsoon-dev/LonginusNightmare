package font

import (
	"fmt"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"golang.org/x/image/font"
)

func NewArcade() (*Arcade, error) {
	ttf, err := truetype.Parse(fonts.ArcadeN_ttf)
	if err != nil {
		return nil, fmt.Errorf("parse ttf: %w", err)
	}

	title := truetype.NewFace(ttf, &truetype.Options{
		Size:    TitleFontSize,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})

	return &Arcade{
		ttf:   ttf,
		title: title,
	}, nil
}

type Arcade struct {
	ttf   *truetype.Font
	title font.Face
}

func (a *Arcade) Title() font.Face {
	return a.title
}

const (
	TitleFontSize = 8
	dpi           = 72
)
