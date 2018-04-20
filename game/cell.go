package game

type Cell struct {
	Alive bool
	Type  int
}

func (c *Cell) Reset() {
	c.Alive = false
	c.Type = 0
}
