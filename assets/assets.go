package assets

import (
	"github.com/eleniums/game-of-life-go/assets/fonts"
	"github.com/eleniums/game-of-life-go/assets/images"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/text"
)

var (
	// Title is a picture of the main logo.
	Title pixel.Picture

	// Icon16x16 is a picture of the 16x16 icon.
	Icon16x16 pixel.Picture

	// CellMap is a picture of the cell sprite map.
	CellMap pixel.Picture

	// GrassMap is a picture of the grass tile sprite map.
	GrassMap pixel.Picture

	// PixelAtlas is the font atlas of the Pixel font.
	PixelAtlas *text.Atlas
)

// Load all assets.
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

	PixelAtlas, err = fonts.Load(fonts.PixelData, 40)
	if err != nil {
		return err
	}

	return nil
}
