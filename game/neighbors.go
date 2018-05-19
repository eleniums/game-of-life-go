package game

// countNeighbors will count the number of living neighbors surrounding a cell.
func (m *Manager) countNeighbors(x, y int, trackDead bool) (count, cross, plus, circle, dot int) {
	var pos Position

	for pos.X = x - 1; pos.X <= x+1; pos.X++ {
		for pos.Y = y - 1; pos.Y <= y+1; pos.Y++ {
			if !(pos.X == x && pos.Y == y) {
				if cell, ok := m.cells[pos]; ok {
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
				} else if trackDead {
					m.dead[pos] = CellTypeCross
				}
			}
		}
	}

	return cross + plus + circle + dot, cross, plus, circle, dot
}
