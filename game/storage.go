package game

import (
	"encoding/json"
	"io/ioutil"
)

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
