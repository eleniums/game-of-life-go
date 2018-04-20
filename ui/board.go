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

	visibleBoardW = 96
	visibleBoardH = 96
)

var (
	SetCellType = game.CellType_Cross
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
		changeCell(cells, win.MousePosition(), true, SetCellType)
	} else if win.JustPressed(pixelgl.MouseButtonRight) {
		changeCell(cells, win.MousePosition(), false, SetCellType)
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
	for x := visibleBoardW; x < game.GridMaxX-visibleBoardW; x++ {
		for y := visibleBoardH; y < game.GridMaxY-visibleBoardH; y++ {
			if cells[x][y].Alive {
				switch cells[x][y].Type {
				case 0:
					draw(b.cellBatch, sprites.Cell1, x-visibleBoardW, y-visibleBoardH)
				case 1:
					draw(b.cellBatch, sprites.Cell2, x-visibleBoardW, y-visibleBoardH)
				case 2:
					draw(b.cellBatch, sprites.Cell3, x-visibleBoardW, y-visibleBoardH)
				case 3:
					draw(b.cellBatch, sprites.Cell4, x-visibleBoardW, y-visibleBoardH)
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

func changeCell(cells game.CellGrid, pos pixel.Vec, alive bool, cellType game.CellType) {
	x := int(pos.X/10 + visibleBoardW)
	y := int(pos.Y/10 + visibleBoardH)

	if x < visibleBoardW || x >= game.GridMaxX-visibleBoardW || y < visibleBoardH || y >= game.GridMaxY-visibleBoardH {
		// do nothing
	} else {
		cells[x][y].Alive = alive
		cells[x][y].Type = cellType
	}
}
