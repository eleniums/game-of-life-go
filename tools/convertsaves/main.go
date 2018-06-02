package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/eleniums/game-of-life-go/game"
)

// Tool for converting an old save file to the new save format.
// Usage: go run main.go <savename>
func main() {
	path := os.Args[1]

	cells, err := load(path)
	if err != nil {
		panic(fmt.Sprintf("unable to load file at %s: %v", path, err))
	}

	err = save(cells, path)
	if err != nil {
		panic(fmt.Sprintf("unable to save file at %s: %v", path, err))
	}
}

type cell struct {
	Alive bool
	Type  game.CellType
}

type storageOld struct {
	Cell *cell
	X    int
	Y    int
}

type storageNew struct {
	Cell game.CellType
	X    int
	Y    int
}

func load(path string) ([]*storageOld, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var compact []*storageOld
	err = json.Unmarshal(data, &compact)

	return compact, err
}

func save(cells []*storageOld, path string) error {
	compact := []*storageNew{}
	for _, v := range cells {
		compact = append(compact, &storageNew{
			Cell: v.Cell.Type,
			X:    v.X,
			Y:    v.Y,
		})
	}

	data, err := json.Marshal(compact)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path, data, 0644)

	return err
}
