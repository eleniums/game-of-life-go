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
	// dimensions of grass tiles
	boardMaxX = 6
	boardMaxY = 6

	// dimensions of visible board space
	visibleBoardW = 96
	visibleBoardH = 96
)

var (
	// SetCellType is the type of cell to place when the grid is clicked.
	SetCellType = game.CellTypeCross

	// lower left corner of visible space
	boardPos = pixel.V(96.0, 96.0)
)

// Board contains batches for drawing the grass and cells.
type Board struct {
	grassGrid  [][]int
	grassBatch *pixel.Batch
	cellBatch  *pixel.Batch
}

// NewBoard creates a new board.
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

// Update the board with any new mouse clicks.
func (b *Board) Update(win *pixelgl.Window, cells game.Grid) {
	if win.JustPressed(pixelgl.MouseButtonLeft) {
		changeCell(cells, win.MousePosition(), true, SetCellType)
	} else if win.JustPressed(pixelgl.MouseButtonRight) {
		changeCell(cells, win.MousePosition(), false, SetCellType)
	}
}

// Draw the board to the screen.
func (b *Board) Draw(t pixel.Target, cells game.Grid) {
	b.drawGrass(t)
	b.drawCells(t, cells)
}

// drawGrass will draw the grass tiles to the board.
func (b *Board) drawGrass(t pixel.Target) {
	b.grassBatch.Clear()

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

	b.grassBatch.Draw(t)
}

// drawCells will draw the cells to the board.
func (b *Board) drawCells(t pixel.Target, cells game.Grid) {
	b.cellBatch.Clear()

	// draw cells to batch
	for k, v := range cells {
		// check if cell is visible and skip if it is not visible
		if k.X < int(boardPos.X) || k.Y < int(boardPos.Y) || k.X >= int(boardPos.X)+visibleBoardW || k.Y >= int(boardPos.Y)+visibleBoardH {
			continue
		}

		switch v {
		case 0:
			draw(b.cellBatch, sprites.Cell1, k.X-visibleBoardW, k.Y-visibleBoardH)
		case 1:
			draw(b.cellBatch, sprites.Cell2, k.X-visibleBoardW, k.Y-visibleBoardH)
		case 2:
			draw(b.cellBatch, sprites.Cell3, k.X-visibleBoardW, k.Y-visibleBoardH)
		case 3:
			draw(b.cellBatch, sprites.Cell4, k.X-visibleBoardW, k.Y-visibleBoardH)
		default:
		}
	}

	b.cellBatch.Draw(t)
}

// draw will draw a single cell to a batch at the given location.
func draw(batch *pixel.Batch, tile *pixel.Sprite, x, y int) {
	loc := pixel.V(tile.Frame().W()/2+tile.Frame().W()*float64(x), tile.Frame().H()/2+tile.Frame().H()*float64(y))
	tile.Draw(batch, pixel.IM.Moved(loc))
}

// changeCell will switch a cell to a different state and type.
func changeCell(cells game.Grid, pos pixel.Vec, alive bool, cellType game.CellType) {
	x := int(pos.X/10 + visibleBoardW)
	y := int(pos.Y/10 + visibleBoardH)

	// check if position is on the board and do nothing if it is out of bounds
	if x < int(boardPos.X) || x >= int(boardPos.X)+visibleBoardW || y < int(boardPos.Y) || y >= int(boardPos.Y)+visibleBoardH {
		return
	}

	// change cell properties
	if alive {
		cells.Add(x, y, cellType)
	} else {
		cells.Delete(x, y)
	}
}
