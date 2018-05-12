package game

import (
	"github.com/eleniums/grid"
)

// countNeighbors will count the number of living neighbors surrounding a cell.
func countNeighbors(g grid.Grid, x, y int) (count, cross, plus, circle, dot int) {
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			if isAlive(g, i, j) && !(i == x && j == y) {
				switch g[i][j].Type {
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
func isAlive(g grid.Grid, x, y int) bool {
	if x < 0 || x >= GridMaxX || y < 0 || y >= GridMaxY {
		return false
	}

	return g[x][y].Alive
}
