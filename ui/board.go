package ui

import (
	"math"
	"math/rand"

	"github.com/eleniums/game-of-life-go/assets"
	"github.com/eleniums/game-of-life-go/game"
	"github.com/eleniums/game-of-life-go/sprites"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

const (
	// dimensions of grass tiles
	boardMaxX = 8
	boardMaxY = 8

	// dimensions of visible board space
	visibleBoardW = 97
	visibleBoardH = 97

	// grid scrolling speed
	scrollSpeed = 20.0
)

var (
	// SetCellType is the type of cell to place when the grid is clicked.
	SetCellType = game.CellTypeCross

	// lower left corner of visible space
	boardPos = pixel.V(0.0, 0.0)
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
func (b *Board) Update(win *pixelgl.Window, dt float64, running bool, cells game.Grid) {
	// add or remove cells using mouse buttons
	if !running {
		if win.JustPressed(pixelgl.MouseButtonLeft) {
			changeCell(cells, win.MousePosition(), true, SetCellType)
		} else if win.JustPressed(pixelgl.MouseButtonRight) {
			changeCell(cells, win.MousePosition(), false, SetCellType)
		}
	}

	// scroll view of grid
	if win.Pressed(pixelgl.KeyLeft) {
		boardPos.X -= scrollSpeed * dt
	}
	if win.Pressed(pixelgl.KeyRight) {
		boardPos.X += scrollSpeed * dt
	}
	if win.Pressed(pixelgl.KeyDown) {
		boardPos.Y -= scrollSpeed * dt
	}
	if win.Pressed(pixelgl.KeyUp) {
		boardPos.Y += scrollSpeed * dt
	}

	// reset view of grid
	if win.Pressed(pixelgl.KeySpace) {
		boardPos.X = 0.0
		boardPos.Y = 0.0
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

	// use values normalized to the size of the board
	normW := boardMaxX * sprites.Grass1.Frame().W()
	normH := boardMaxY * sprites.Grass1.Frame().H()

	normX := math.Mod(boardPos.X, normW)
	if normX < 0 {
		normX += normW
	}

	normY := math.Mod(boardPos.Y, normH)
	if normY < 0 {
		normY += normH
	}

	// get starting grassGrid coordinates
	tileX := int(normX) / 16 % boardMaxX
	tileY := int(normY) / 16 % boardMaxY

	for i := 0; i < boardMaxX; i++ {
		for j := 0; j < boardMaxY; j++ {

			xpos := float64(i-1) - math.Mod(normX, 16)/16
			ypos := float64(j-1) - math.Mod(normY, 16)/16

			// draw grass to batch
			switch b.grassGrid[tileX][tileY] {
			case 0:
				draw(b.grassBatch, sprites.Grass1, xpos, ypos)
			case 1:
				draw(b.grassBatch, sprites.Grass2, xpos, ypos)
			case 2:
				draw(b.grassBatch, sprites.Grass3, xpos, ypos)
			case 3:
				draw(b.grassBatch, sprites.Grass4, xpos, ypos)
			default:
			}

			tileY++
			if tileY >= boardMaxY {
				tileY = 0
			}
		}

		tileX++
		if tileX >= boardMaxX {
			tileX = 0
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

		xpos := float64(k.X) - boardPos.X
		ypos := float64(k.Y) - boardPos.Y

		switch v {
		case game.CellTypeCross:
			draw(b.cellBatch, sprites.Cell1, xpos, ypos)
		case game.CellTypePlus:
			draw(b.cellBatch, sprites.Cell2, xpos, ypos)
		case game.CellTypeCircle:
			draw(b.cellBatch, sprites.Cell3, xpos, ypos)
		case game.CellTypeDot:
			draw(b.cellBatch, sprites.Cell4, xpos, ypos)
		default:
		}
	}

	b.cellBatch.Draw(t)
}

// draw will draw a single cell to a batch at the given location.
func draw(batch *pixel.Batch, sprite *pixel.Sprite, x, y float64) {
	loc := pixel.V(sprite.Frame().W()/2+sprite.Frame().W()*x, sprite.Frame().H()/2+sprite.Frame().H()*y)
	sprite.Draw(batch, pixel.IM.Moved(loc))
}

// changeCell will switch a cell to a different state and type.
func changeCell(cells game.Grid, pos pixel.Vec, alive bool, cellType game.CellType) {
	x := int(pos.X/10 + boardPos.X)
	y := int(pos.Y/10 + boardPos.Y)

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
