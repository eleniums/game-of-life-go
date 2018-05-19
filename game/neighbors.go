package game

// countNeighbors will count the number of living neighbors surrounding a cell.
func (m *Manager) countNeighbors(x, y int, trackDead bool) (int, []int) {
	types := make([]int, 4)
	count := 0

	var pos Position
	for pos.X = x - 1; pos.X <= x+1; pos.X++ {
		for pos.Y = y - 1; pos.Y <= y+1; pos.Y++ {
			if !(pos.X == x && pos.Y == y) {
				if cell, ok := m.cells[pos]; ok {
					count++
					types[cell]++
				} else if trackDead {
					m.dead[pos] = CellTypeCross
				}
			}
		}
	}

	return count, types
}
