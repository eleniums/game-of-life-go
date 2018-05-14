package game

import (
	"encoding/json"
	"io/ioutil"

	"github.com/eleniums/game-of-life-go/game"
)

// storage defines the type used to store a cell in a file.
type storage struct {
	Cell *Cell
	X    int
	Y    int
}

// save a grid to a file.
func save(cells game.Grid, path string) error {
	compact := []*storage{}
	for k, v := range cells {
		compact = append(compact, &storage{
			Cell: &Cell{
				Alive: true,
				Type:  v,
			},
			X: k.X,
			Y: k.Y,
		})
	}

	data, err := json.Marshal(compact)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path, data, 0644)

	return err
}

// load a pattern from a file to a grid.
func load(path string) (game.Grid, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var compact []*storage
	err = json.Unmarshal(data, &compact)

	cells := game.NewGrid()
	for _, v := range compact {
		cells.Add(v.X, v.Y, v.Cell.Type)
	}

	return cells, err
}
