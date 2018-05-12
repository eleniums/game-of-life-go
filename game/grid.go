package game

const (
	// GridMaxX is the width of the grid.
	GridMaxX = 288

	// GridMaxY is the height of the grid.
	GridMaxY = 288
)

// CellGrid is a grid of cells.
type CellGrid [][]*Cell

// NewCellGrid will create a new grid with default values.
func NewCellGrid() CellGrid {
	grid := make(CellGrid, GridMaxX)
	for x := 0; x < GridMaxX; x++ {
		grid[x] = make([]*Cell, GridMaxY)
		for y := 0; y < GridMaxY; y++ {
			grid[x][y] = &Cell{
				Alive: false,
				Type:  0,
			}
		}
	}

	return grid
}

// Clear will reset all cells in the grid to their default values.
func (c CellGrid) Clear() {
	for x := range c {
		for y := range c[x] {
			c[x][y].Clear()
		}
	}
}

// Copy will make a copy of the grid.
func (c CellGrid) Copy() CellGrid {
	grid := make(CellGrid, GridMaxX)
	for x := 0; x < GridMaxX; x++ {
		grid[x] = make([]*Cell, GridMaxY)
		for y := 0; y < GridMaxY; y++ {
			grid[x][y] = c[x][y].Copy()
		}
	}

	return grid
}
