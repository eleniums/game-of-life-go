package sprites

import (
	"github.com/eleniums/game-of-life-go/assets"
	"github.com/faiface/pixel"
)

var (
	Title     *pixel.Sprite
	Icon16x16 *pixel.Sprite

	Cell1 *pixel.Sprite
	Cell2 *pixel.Sprite
	Cell3 *pixel.Sprite
	Cell4 *pixel.Sprite

	Grass1 *pixel.Sprite
	Grass2 *pixel.Sprite
	Grass3 *pixel.Sprite
	Grass4 *pixel.Sprite
)

func Load() {
	Title = pixel.NewSprite(assets.Title, assets.Title.Bounds())
	Icon16x16 = pixel.NewSprite(assets.Icon16x16, assets.Icon16x16.Bounds())

	Cell1 = pixel.NewSprite(assets.CellMap, pixel.R(0, 10, 10, 20))
	Cell2 = pixel.NewSprite(assets.CellMap, pixel.R(10, 10, 20, 20))
	Cell3 = pixel.NewSprite(assets.CellMap, pixel.R(0, 0, 10, 10))
	Cell4 = pixel.NewSprite(assets.CellMap, pixel.R(10, 0, 20, 10))

	Grass1 = pixel.NewSprite(assets.GrassMap, pixel.R(0, 0, 160, 160))
	Grass2 = pixel.NewSprite(assets.GrassMap, pixel.R(160, 0, 320, 160))
	Grass3 = pixel.NewSprite(assets.GrassMap, pixel.R(0, 160, 160, 320))
	Grass4 = pixel.NewSprite(assets.GrassMap, pixel.R(160, 160, 320, 320))
}
