package game

import (
	"github.com/eleniums/grid"
)

// countNeighbors will count the number of living neighbors surrounding a cell.
func countNeighbors(cells grid.Grid, x, y int) (count, cross, plus, circle, dot int) {
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			if isAlive(cells, i, j) && !(i == x && j == y) {
				switch cells[i][j].Type {
				case CellTypeCross:
					cross++
				case CellTypePlus:
					plus++
				case CellTypeCircle:
					circle++
				case CellTypeDot:
					dot++
				}
			}
		}
	}

	return cross + plus + circle + dot, cross, plus, circle, dot
}

// isAlive will determine if the cell at the given position is alive.
func isAlive(cells grid.Grid, x, y int) bool {
	cell, ok := cells.Retrieve(float64(x), float64(y))
	return ok && cell.(*Cell).Alive
}
