package game

func applyRules(alive bool, neighbors int) bool {
	if alive {
		return neighbors >= 2 && neighbors <= 3
	} else {
		return neighbors == 3
	}
}
