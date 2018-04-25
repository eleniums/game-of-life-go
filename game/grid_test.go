package game

import (
	"testing"
)

func Benchmark_CellGrid_CountNeighbors_BestCase(b *testing.B) {
	cells := NewCellGrid()

	cells[4][5].Alive = false
	cells[6][5].Alive = false
	cells[5][6].Alive = false
	cells[4][6].Alive = false
	cells[6][6].Alive = false
	cells[5][4].Alive = false
	cells[4][4].Alive = false
	cells[6][4].Alive = false

	for i := 0; i < b.N; i++ {
		cells.CountNeighbors(5, 5)
	}
}

func Benchmark_CellGrid_CountNeighbors_WorstCase(b *testing.B) {
	cells := NewCellGrid()

	cells[4][5].Alive = true
	cells[6][5].Alive = true
	cells[5][6].Alive = true
	cells[4][6].Alive = true
	cells[6][6].Alive = true
	cells[5][4].Alive = true
	cells[4][4].Alive = true
	cells[6][4].Alive = true

	for i := 0; i < b.N; i++ {
		cells.CountNeighbors(5, 5)
	}
}
