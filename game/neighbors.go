package game

// countNeighbors will count the number of living neighbors surrounding a cell.
func countNeighbors(cells Grid, x, y int) (count, cross, plus, circle, dot int) {
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			if !(i == x && j == y) {
				if cell, ok := cells.Retrieve(x, y); ok {
					switch cell {
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
	}

	return cross + plus + circle + dot, cross, plus, circle, dot
}
