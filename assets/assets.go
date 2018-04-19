package assets

import (
	"github.com/eleniums/game-of-life-go/assets/fonts"
	"github.com/eleniums/game-of-life-go/assets/images"
	"github.com/faiface/pixel"
	"golang.org/x/image/font"
)

var (
	Icon16x16 pixel.Picture
	Title     pixel.Picture
	CellMap   pixel.Picture
	GrassMap  pixel.Picture
	PixelFont font.Face
)

func Load() error {
	var err error

	Icon16x16, err = images.Load(images.Icon16x16Data)
	if err != nil {
		return err
	}

	Title, err = images.Load(images.TitleData)
	if err != nil {
		return err
	}

	CellMap, err = images.Load(images.CellMapData)
	if err != nil {
		return err
	}

	GrassMap, err = images.Load(images.GrassMapData)
	if err != nil {
		return err
	}

	PixelFont, err = fonts.Load(fonts.PixelData, 40)
	if err != nil {
		return err
	}

	return nil
}
