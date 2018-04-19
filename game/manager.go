package game

type Manager struct {
	cells   CellGrid
	buffer  CellGrid
	running bool
}

func NewManager() *Manager {
	return &Manager{
		cells:   NewCellGrid(GridMaxX, GridMaxY),
		buffer:  NewCellGrid(GridMaxX, GridMaxY),
		running: false,
	}
}

func (m *Manager) Update() {
	// iterate over grid and apply rules
	for x := range m.cells {
		for y := range m.cells[x] {
			neighbors := countNeighbors(m.cells, x, y)
			m.buffer[x][y].Alive = applyRules(m.cells[x][y].Alive, neighbors)
		}
	}

	// swap active cells with buffer
	temp := m.cells
	m.cells = m.buffer
	m.buffer = temp
}

func (m *Manager) Cells() CellGrid {
	return m.cells
}

func (m *Manager) Running() bool {
	return m.running
}
