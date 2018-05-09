package game

// CellType is the design and color of the cell.
type CellType int

const (
	// CellTypeCross has a cross pattern in the middle.
	CellTypeCross CellType = iota

	// CellTypePlus has a plus pattern in the middle.
	CellTypePlus

	// CellTypeCircle has a circle pattern in the middle.
	CellTypeCircle

	// CellTypeDot has a dot pattern in the middle.
	CellTypeDot
)

// Cell represents a point that is either alive or dead.
type Cell struct {
	Alive bool
	Type  CellType
}

// Clear will reset the cell to the defaults.
func (c *Cell) Clear() {
	c.Alive = false
	c.Type = CellTypeCross
}

// Copy will make a perfect copy of the cell.
func (c *Cell) Copy() *Cell {
	return &Cell{
		Alive: c.Alive,
		Type:  c.Type,
	}
}
