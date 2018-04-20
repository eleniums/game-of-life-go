package game

import "time"

var (
	Interval = 1000
)

type Manager struct {
	cells   CellGrid
	buffer  CellGrid
	running bool
	ticker  *time.Ticker
}

func NewManager() *Manager {
	return &Manager{
		cells:   NewCellGrid(GridMaxX, GridMaxY),
		buffer:  NewCellGrid(GridMaxX, GridMaxY),
		running: false,
	}
}

func (m *Manager) Update() {
	if m.running {
		select {
		case <-m.ticker.C:
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
		default:
		}
	}
}

func (m *Manager) Start() {
	m.running = true
	m.ticker = time.NewTicker(time.Duration(Interval) * time.Millisecond)
}

func (m *Manager) Stop() {
	m.ticker.Stop()
	m.running = false
}

func (m *Manager) Cells() CellGrid {
	return m.cells
}

func (m *Manager) Running() bool {
	return m.running
}
