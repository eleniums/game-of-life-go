package game

import (
	"time"
)

var (
	Interval = 100
)

type Manager struct {
	cells   CellGrid
	buffer  CellGrid
	memory  CellGrid
	running bool
	ticker  *time.Ticker
}

func NewManager() *Manager {
	return &Manager{
		cells:   NewCellGrid(),
		buffer:  NewCellGrid(),
		memory:  NewCellGrid(),
		running: false,
	}
}

func (m *Manager) Update() {
	if m.running {
		select {
		case <-m.ticker.C:
			m.updateBuffer()
			m.swapBuffer()
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

func (m *Manager) Clear() {
	m.cells.Clear()
	m.buffer.Clear()
}

func (m *Manager) Store() {
	m.memory = m.cells.Copy()
}

func (m *Manager) Reset() {
	m.cells = m.memory.Copy()
}

func (m *Manager) Save(path string) error {
	return save(m.cells, path)
}

func (m *Manager) Load(path string) error {
	cells, err := load(path)
	if err != nil {
		return err
	}

	m.cells = cells
	m.buffer.Clear()
	m.memory = m.cells.Copy()

	return nil
}

func (m *Manager) Cells() CellGrid {
	return m.cells
}

func (m *Manager) Running() bool {
	return m.running
}

// updateBuffer will iterate over the grid and apply rules.
func (m *Manager) updateBuffer() {
	for x := range m.cells {
		for y := range m.cells[x] {
			neighbors, cross, plus, circle, dot := m.cells.CountNeighbors(x, y)
			if m.cells[x][y].Alive {
				m.buffer[x][y].Type = m.cells[x][y].Type
			} else if neighbors == 3 {
				m.buffer[x][y].Type = determineType(cross, plus, circle, dot)
			}
			m.buffer[x][y].Alive = applyRules(m.cells[x][y].Alive, neighbors)
		}
	}
}

// swapBuffer will swap active cells with buffer
func (m *Manager) swapBuffer() {
	temp := m.cells
	m.cells = m.buffer
	m.buffer = temp
}
