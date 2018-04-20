package game

const (
	GridMaxX = 96
	GridMaxY = 96
)

type CellGrid [][]*Cell

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

func (c CellGrid) Clear() {
	for x := range c {
		for y := range c[x] {
			c[x][y].Clear()
		}
	}
}

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
