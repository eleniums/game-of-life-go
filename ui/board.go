package ui

import (
	"log"
	"math"
	"math/rand"

	"github.com/eleniums/game-of-life-go/assets"
	"github.com/eleniums/game-of-life-go/game"
	"github.com/eleniums/game-of-life-go/sprites"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

const (
	// default dimensions of grass tiles
	defaultGrassW = 8
	defaultGrassH = 8

	// grid scrolling speed
	scrollSpeed = 20.0
)

var (
	// SetCellType is the type of cell to place when the grid is clicked.
	SetCellType = game.CellTypeCross
)

// Board contains batches for drawing the grass and cells.
type Board struct {
	grassGrid  [][]int
	grassBatch *pixel.Batch
	cellBatch  *pixel.Batch

	// dimensions of grass tiles
	grassW int
	grassH int

	// dimensions of visible cells
	cellW int
	cellH int

	// lower left corner of visible space
	pos pixel.Vec
}

// NewBoard creates a new board.
func NewBoard() *Board {
	grassBatch := pixel.NewBatch(&pixel.TrianglesData{}, assets.GrassMap)
	cellBatch := pixel.NewBatch(&pixel.TrianglesData{}, assets.CellMap)

	// randomize background tiles
	grassGrid := make([][]int, defaultGrassW)
	for x := 0; x < defaultGrassW; x++ {
		grassGrid[x] = make([]int, defaultGrassH)
		for y := 0; y < defaultGrassH; y++ {
			grassGrid[x][y] = rand.Intn(4)
		}
	}

	return &Board{
		grassGrid:  grassGrid,
		grassBatch: grassBatch,
		cellBatch:  cellBatch,
		pos:        pixel.V(0.0, 0.0),
	}
}

// Update the board with any new mouse clicks.
func (b *Board) Update(win *pixelgl.Window, dt float64, running bool, cells game.Grid) {
	// add or remove cells using mouse buttons
	if !running {
		if win.JustPressed(pixelgl.MouseButtonLeft) {
			b.changeCell(cells, win.MousePosition(), true, SetCellType)
		} else if win.JustPressed(pixelgl.MouseButtonRight) {
			b.changeCell(cells, win.MousePosition(), false, SetCellType)
		}
	}

	// scroll view of grid
	if win.Pressed(pixelgl.KeyLeft) {
		b.pos.X -= scrollSpeed * dt
	}
	if win.Pressed(pixelgl.KeyRight) {
		b.pos.X += scrollSpeed * dt
	}
	if win.Pressed(pixelgl.KeyDown) {
		b.pos.Y -= scrollSpeed * dt
	}
	if win.Pressed(pixelgl.KeyUp) {
		b.pos.Y += scrollSpeed * dt
	}

	// reset view of grid
	if win.Pressed(pixelgl.KeySpace) {
		b.pos.X = 0.0
		b.pos.Y = 0.0
	}
}

// Draw the board to the screen.
func (b *Board) Draw(t pixel.Target, cells game.Grid) {
	b.drawGrass(t)
	b.drawCells(t, cells)
}

// Resize the board with new dimensions.
func (b *Board) Resize(w, h float64) {

	log.Printf("Resizing to w: %v, h: %v\n", w, h)

	// expand viewable area for cells
	b.cellW = int(math.Ceil((w-300)/sprites.Cell1.Frame().W())) + 1
	b.cellH = int(math.Ceil(h/sprites.Cell1.Frame().H())) + 1

	// expand viewable area for grass
	b.grassW = int(math.Ceil(((w-300)/sprites.Grass1.Frame().W()+2)/defaultGrassW)) * defaultGrassW
	b.grassH = int(math.Ceil((h/sprites.Grass1.Frame().H()+2)/defaultGrassH)) * defaultGrassH

	// expand grid size
	grassGrid := make([][]int, b.grassW)
	for x := 0; x < b.grassW; x++ {
		grassGrid[x] = make([]int, b.grassH)
		for y := 0; y < b.grassH; y++ {
			grassGrid[x][y] = b.grassGrid[x%defaultGrassW][y%defaultGrassH]
		}
	}
	b.grassGrid = grassGrid
}

// drawGrass will draw the grass tiles to the board.
func (b *Board) drawGrass(t pixel.Target) {
	b.grassBatch.Clear()

	// use values normalized to the size of the board
	normW := float64(b.grassW) * sprites.Grass1.Frame().W()
	normH := float64(b.grassH) * sprites.Grass1.Frame().H()

	normX := math.Mod(b.pos.X, normW)
	if normX < 0 {
		normX += normW
	}

	normY := math.Mod(b.pos.Y, normH)
	if normY < 0 {
		normY += normH
	}

	// get starting grassGrid coordinates
	tileX := int(normX) / 16 % b.grassW
	tileY := int(normY) / 16 % b.grassH

	for i := 0; i < b.grassW; i++ {
		for j := 0; j < b.grassH; j++ {

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
			if tileY >= b.grassH {
				tileY = 0
			}
		}

		tileX++
		if tileX >= b.grassW {
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
		if k.X < int(b.pos.X) || k.Y < int(b.pos.Y) || k.X >= int(b.pos.X)+b.cellW || k.Y >= int(b.pos.Y)+b.cellH {
			continue
		}

		xpos := float64(k.X) - b.pos.X
		ypos := float64(k.Y) - b.pos.Y

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
func (b *Board) changeCell(cells game.Grid, pos pixel.Vec, alive bool, cellType game.CellType) {
	x := int(pos.X/10 + b.pos.X)
	y := int(pos.Y/10 + b.pos.Y)

	// check if position is on the board and do nothing if it is out of bounds
	if x < int(b.pos.X) || x >= int(b.pos.X)+b.cellW || y < int(b.pos.Y) || y >= int(b.pos.Y)+b.cellH {
		return
	}

	// change cell properties
	if alive {
		cells.Add(x, y, cellType)
	} else {
		cells.Delete(x, y)
	}
}
