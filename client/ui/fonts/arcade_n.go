package fonts

import (
	"fmt"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"

	"github.com/libmonsoon-dev/LonginusNightmare/client/assets/fonts"
)

const (
	TitleFontSize = 8
	dpi           = 72
)

func NewArcadeN() (*Arcade, error) {
	ttf, err := truetype.Parse(fonts.ArcadeN)
	if err != nil {
		return nil, fmt.Errorf("parse ttf: %w", err)
	}

	title := truetype.NewFace(ttf, &truetype.Options{
		Size:    TitleFontSize,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})

	f := &Arcade{
		title: title,
	}

	return f, nil
}

type Arcade struct {
	title font.Face
}

func (a *Arcade) Title() font.Face {
	return a.title
}
