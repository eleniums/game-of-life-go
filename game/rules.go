package game

// applyRules will determine if a cell should live or die, based on the following rules:
// - If there are less than 2 living cells surrounding a living cell, it will die, as if by underpopulation.
// - If there are 2 or 3 living cells surrounding a living cell, it will continue to live.
// - If there are more than 3 living cells surrounding a living cell, it will die, as if by overpopulation.
// - If there are exactly 3 living cells surrounding a dead cell, it will become alive, as if by reproduction.
func applyRules(alive bool, neighbors int) bool {
	if alive {
		return neighbors >= 2 && neighbors <= 3
	} else {
		return neighbors == 3
	}
}
