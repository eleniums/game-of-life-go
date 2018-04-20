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
				Type:  0,
			}
		}
	}

	return grid
}

func (c CellGrid) IsAlive(x, y int) bool {
	if x < 0 || x >= GridMaxX || y < 0 || y >= GridMaxY {
		return false
	}

	return c[x][y].Alive
}

func (c CellGrid) CountNeighbors(x, y int) int {
	count := 0

	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			if c.IsAlive(i, j) && !(i == x && j == y) {
				count++
			}
		}
	}

	return count
}
