package game

const (
	GridMaxX = 288
	GridMaxY = 288
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

func (c CellGrid) CountNeighbors(x, y int) (count, cross, plus, circle, dot int) {
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			if c.IsAlive(i, j) && !(i == x && j == y) {
				switch c[i][j].Type {
				case CellType_Cross:
					cross++
				case CellType_Plus:
					plus++
				case CellType_Circle:
					circle++
				case CellType_Dot:
					dot++
				}
			}
		}
	}

	return cross + plus + circle + dot, cross, plus, circle, dot
}
