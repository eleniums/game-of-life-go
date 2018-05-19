package game

import (
	"time"
)

var (
	// Interval is the time between grid updates in milliseconds.
	Interval = 100
)

// Manager controls the grid updates.
type Manager struct {
	cells   Grid
	buffer  Grid
	memory  Grid
	dead    Grid
	running bool
	ticker  *time.Ticker
}

// NewManager creates a new Manager.
func NewManager() *Manager {
	return &Manager{
		cells:   NewGrid(),
		buffer:  NewGrid(),
		memory:  NewGrid(),
		dead:    NewGrid(),
		running: false,
	}
}

// Update the grid buffer and swap it with the currently displayed grid if the interval has elapsed.
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

// Start the grid updates.
func (m *Manager) Start() {
	m.running = true
	m.ticker = time.NewTicker(time.Duration(Interval) * time.Millisecond)
}

// Stop the grid updates.
func (m *Manager) Stop() {
	m.ticker.Stop()
	m.running = false
}

// Clear the grid.
func (m *Manager) Clear() {
	m.cells.Clear()
	m.buffer.Clear()
}

// Store the grid in memory.
func (m *Manager) Store() {
	m.memory = m.cells.Copy()
}

// Reset the grid to the copy stored in memory.
func (m *Manager) Reset() {
	m.cells = m.memory.Copy()
}

// Save the grid to a file.
func (m *Manager) Save(path string) error {
	return save(m.cells, path)
}

// Load a pattern from a file into the grid.
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

// Cells returns the grid.
func (m *Manager) Cells() Grid {
	return m.cells
}

// Running returns true if the simulation is active.
func (m *Manager) Running() bool {
	return m.running
}

// updateBuffer will iterate over the grid and apply rules.
func (m *Manager) updateBuffer() {
	m.buffer = NewGrid()
	m.dead = NewGrid()

	for k, v := range m.cells {
		neighbors, _, _, _, _ := m.countNeighbors(k.X, k.Y, true)
		if applyRules(true, neighbors) {
			m.buffer[k] = v
		}
	}

	for k := range m.dead {
		neighbors, cross, plus, circle, dot := m.countNeighbors(k.X, k.Y, false)
		if applyRules(false, neighbors) {
			m.buffer[k] = determineType(cross, plus, circle, dot)
		}
	}
}

// swapBuffer will swap active cells with buffer.
func (m *Manager) swapBuffer() {
	m.cells = m.buffer
}
