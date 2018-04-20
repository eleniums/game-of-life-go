package game

type Cell struct {
	Alive bool
	Type  int
}

func (c *Cell) Clear() {
	c.Alive = false
	c.Type = 0
}

func (c *Cell) Copy() *Cell {
	return &Cell{
		Alive: c.Alive,
		Type:  c.Type,
	}
}
