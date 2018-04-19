package ui

import (
	"math/rand"

	"github.com/eleniums/game-of-life-go/assets"
	"github.com/eleniums/game-of-life-go/sprites"
	"github.com/faiface/pixel"
)

const (
	boardMaxX = 6
	boardMaxY = 6
)

type Board struct {
	batch *pixel.Batch
	grid  [][]int
}

func NewBoard() *Board {
	batch := pixel.NewBatch(&pixel.TrianglesData{}, assets.GrassMap)

	// randomize background tiles
	grid := make([][]int, boardMaxX)
	for x := 0; x < boardMaxX; x++ {
		grid[x] = make([]int, boardMaxY)
		for y := 0; y < boardMaxY; y++ {
			grid[x][y] = rand.Intn(4)
		}
	}

	return &Board{
		batch: batch,
		grid:  grid,
	}
}

func (b *Board) Draw(t pixel.Target) {
	b.batch.Clear()

	for x := range b.grid {
		for y := range b.grid[x] {
			switch b.grid[x][y] {
			case 0:
				draw(b.batch, sprites.Grass1, x, y)
			case 1:
				draw(b.batch, sprites.Grass2, x, y)
			case 2:
				draw(b.batch, sprites.Grass3, x, y)
			case 3:
				draw(b.batch, sprites.Grass4, x, y)
			default:
			}
		}
	}

	b.batch.Draw(t)
}

func draw(batch *pixel.Batch, tile *pixel.Sprite, xpos, ypos int) {
	loc := pixel.V(tile.Frame().W()/2+tile.Frame().W()*float64(xpos), tile.Frame().H()/2+tile.Frame().H()*float64(ypos))
	tile.Draw(batch, pixel.IM.Moved(loc))
}
