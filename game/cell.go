package game

type CellType int

const (
	CellType_Cross CellType = iota
	CellType_Plus
	CellType_Circle
	CellType_Dot
)

type Cell struct {
	Alive bool
	Type  CellType
}

func (c *Cell) Clear() {
	c.Alive = false
	c.Type = CellType_Cross
}

func (c *Cell) Copy() *Cell {
	return &Cell{
		Alive: c.Alive,
		Type:  c.Type,
	}
}
