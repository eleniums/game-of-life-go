package game

type CellType int

const (
	CellTypeCross CellType = iota
	CellTypePlus
	CellTypeCircle
	CellTypeDot
)

type Cell struct {
	Alive bool
	Type  CellType
}

func (c *Cell) Clear() {
	c.Alive = false
	c.Type = CellTypeCross
}

func (c *Cell) Copy() *Cell {
	return &Cell{
		Alive: c.Alive,
		Type:  c.Type,
	}
}
