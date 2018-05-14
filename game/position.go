package game

// Position is a 2D point in space.
type Position struct {
	X int
	Y int
}

// NewPosition will create a new position.
func NewPosition(x, y int) Position {
	return Position{
		X: x,
		Y: y,
	}
}
