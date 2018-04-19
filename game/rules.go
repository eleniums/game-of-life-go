package game

func applyRules(alive bool, neighbors int) bool {
	if alive {
		return neighbors >= 2 && neighbors <= 3
	} else {
		return neighbors == 3
	}
}

func countNeighbors(cells CellGrid, x, y int) int {
	count := 0

	for i := x - 1; i < x+1 && i < GridMaxX; i++ {
		if i < 0 {
			i = 0
		}
		for j := y - 1; j < y+1 && j < GridMaxY; j++ {
			if j < 0 {
				j = 0
			}
			if cells[i][j].Alive && !(i == x && j == y) {
				count++
			}
		}
	}

	return count
}
