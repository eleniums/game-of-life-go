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
