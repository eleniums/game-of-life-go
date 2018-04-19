package game

const (
	GridMaxX = 96
	GridMaxY = 96
)

type CellGrid [][]*Cell

func NewCellGrid(xdim, ydim int) CellGrid {
	grid := make(CellGrid, xdim)
	for x := 0; x < xdim; x++ {
		grid[x] = make([]*Cell, ydim)
		for y := 0; y < ydim; y++ {
			grid[x][y] = &Cell{
				Alive: false,
				Type: 0,
			}
		}
	}

	return grid
}
