package game

import (
	"time"
)

var (
	Interval = 1000
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
			// iterate over grid and apply rules
			for x := range m.cells {
				for y := range m.cells[x] {
					neighbors := m.cells.CountNeighbors(x, y)
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

func (m *Manager) Cells() CellGrid {
	return m.cells
}

func (m *Manager) Running() bool {
	return m.running
}
