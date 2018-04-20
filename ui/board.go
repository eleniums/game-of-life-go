package ui

import (
	"math/rand"

	"github.com/eleniums/game-of-life-go/assets"
	"github.com/eleniums/game-of-life-go/game"
	"github.com/eleniums/game-of-life-go/sprites"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

const (
	boardMaxX = 6
	boardMaxY = 6
)

type Board struct {
	grassGrid  [][]int
	grassBatch *pixel.Batch
	cellBatch  *pixel.Batch
}

func NewBoard() *Board {
	grassBatch := pixel.NewBatch(&pixel.TrianglesData{}, assets.GrassMap)
	cellBatch := pixel.NewBatch(&pixel.TrianglesData{}, assets.CellMap)

	// randomize background tiles
	grassGrid := make([][]int, boardMaxX)
	for x := 0; x < boardMaxX; x++ {
		grassGrid[x] = make([]int, boardMaxY)
		for y := 0; y < boardMaxY; y++ {
			grassGrid[x][y] = rand.Intn(4)
		}
	}

	return &Board{
		grassGrid:  grassGrid,
		grassBatch: grassBatch,
		cellBatch:  cellBatch,
	}
}

func (b *Board) Update(win *pixelgl.Window, cells game.CellGrid) {
	if win.JustPressed(pixelgl.MouseButtonLeft) {
		mpos := win.MousePosition()

		x := int(mpos.X / 10)
		y := int(mpos.Y / 10)

		if x < 0 || x >= game.GridMaxX || y < 0 || y >= game.GridMaxY {
			// do nothing
		} else {
			cells[x][y].Alive = !cells[x][y].Alive
		}
	}
}

func (b *Board) Draw(t pixel.Target, cells game.CellGrid) {
	b.grassBatch.Clear()
	b.cellBatch.Clear()

	// draw grass to batch
	for x := range b.grassGrid {
		for y := range b.grassGrid[x] {
			switch b.grassGrid[x][y] {
			case 0:
				draw(b.grassBatch, sprites.Grass1, x, y)
			case 1:
				draw(b.grassBatch, sprites.Grass2, x, y)
			case 2:
				draw(b.grassBatch, sprites.Grass3, x, y)
			case 3:
				draw(b.grassBatch, sprites.Grass4, x, y)
			default:
			}
		}
	}

	// draw cells to batch
	for x := range cells {
		for y := range cells[x] {
			if cells[x][y].Alive {
				switch cells[x][y].Type {
				case 0:
					draw(b.cellBatch, sprites.Cell1, x, y)
				case 1:
					draw(b.cellBatch, sprites.Cell2, x, y)
				case 2:
					draw(b.cellBatch, sprites.Cell3, x, y)
				case 3:
					draw(b.cellBatch, sprites.Cell4, x, y)
				default:
				}
			}
		}
	}

	b.grassBatch.Draw(t)
	b.cellBatch.Draw(t)
}

func draw(batch *pixel.Batch, tile *pixel.Sprite, x, y int) {
	loc := pixel.V(tile.Frame().W()/2+tile.Frame().W()*float64(x), tile.Frame().H()/2+tile.Frame().H()*float64(y))
	tile.Draw(batch, pixel.IM.Moved(loc))
}
