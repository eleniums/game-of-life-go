package game

// countNeighbors will count the number of living neighbors surrounding a cell.
func (m *Manager) countNeighbors(x, y int, dead func(x, y int)) (count, cross, plus, circle, dot int) {
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			if !(i == x && j == y) {
				if cell, ok := m.cells.Retrieve(i, j); ok {
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
				} else if dead != nil {
					dead(i, j)
				}
			}
		}
	}

	return cross + plus + circle + dot, cross, plus, circle, dot
}
