package game

import (
	"encoding/json"
	"io/ioutil"

	"github.com/eleniums/grid"
)

// storage defines the type used to store a cell in a file.
type storage struct {
	Cell *Cell
	X    int
	Y    int
}

// save a grid to a file.
func save(cells grid.Grid, path string) error {
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

// load a pattern from a file to a grid.
func load(path string) (grid.Grid, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var compact []*storage
	err = json.Unmarshal(data, &compact)

	cells := grid.NewGrid()
	for _, v := range compact {
		cells[v.X][v.Y] = v.Cell
	}

	return cells, err
}
