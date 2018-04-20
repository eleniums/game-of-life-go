package game

import (
	"encoding/json"
	"io/ioutil"
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
			// iterate over grid and apply rules
			for x := range m.cells {
				for y := range m.cells[x] {
					neighbors, cross, plus, circle, dot := m.cells.CountNeighbors(x, y)
					if m.cells[x][y].Alive {
						m.buffer[x][y].Type = m.cells[x][y].Type
					} else {
						m.buffer[x][y].Type = determineType(cross, plus, circle, dot)
					}
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

	return nil
}

func (m *Manager) Cells() CellGrid {
	return m.cells
}

func (m *Manager) Running() bool {
	return m.running
}

type storage struct {
	Cell *Cell
	X    int
	Y    int
}

func save(cells CellGrid, path string) error {
	compact := []*storage{}
	for x := range cells {
		for y := range cells[x] {
			if cells[x][y].Alive {
				compact = append(compact, &storage{
					Cell: cells[x][y],
					X:    x,
					Y:    y,
				})
			}
		}
	}

	data, err := json.Marshal(compact)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path, data, 0644)

	return err
}

func load(path string) (CellGrid, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var compact []*storage
	err = json.Unmarshal(data, &compact)

	cells := NewCellGrid()
	for _, v := range compact {
		cells[v.X][v.Y] = v.Cell
	}

	return cells, err
}
