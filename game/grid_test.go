package game

import (
	"testing"

	assert "github.com/stretchr/testify/require"
)

func Test_CellGrid_CountNeighbors_Alive(t *testing.T) {

	t.Run("Upper Left", func(t *testing.T) {
		cells := NewCellGrid()
		cells[4][6].Alive = true
		count, _, _, _, _ := cells.CountNeighbors(5, 5)
		assert.Equal(t, 1, count)
	})

	t.Run("Upper Middle", func(t *testing.T) {
		cells := NewCellGrid()
		cells[5][6].Alive = true
		count, _, _, _, _ := cells.CountNeighbors(5, 5)
		assert.Equal(t, 1, count)
	})

	t.Run("Upper Right", func(t *testing.T) {
		cells := NewCellGrid()
		cells[6][6].Alive = true
		count, _, _, _, _ := cells.CountNeighbors(5, 5)
		assert.Equal(t, 1, count)
	})

	t.Run("Middle Left", func(t *testing.T) {
		cells := NewCellGrid()
		cells[4][5].Alive = true
		count, _, _, _, _ := cells.CountNeighbors(5, 5)
		assert.Equal(t, 1, count)
	})

	t.Run("Center", func(t *testing.T) {
		cells := NewCellGrid()
		cells[5][5].Alive = true
		count, _, _, _, _ := cells.CountNeighbors(5, 5)
		assert.Equal(t, 0, count)
	})

	t.Run("Middle Right", func(t *testing.T) {
		cells := NewCellGrid()
		cells[6][5].Alive = true
		count, _, _, _, _ := cells.CountNeighbors(5, 5)
		assert.Equal(t, 1, count)
	})

	t.Run("Lower Left", func(t *testing.T) {
		cells := NewCellGrid()
		cells[4][4].Alive = true
		count, _, _, _, _ := cells.CountNeighbors(5, 5)
		assert.Equal(t, 1, count)
	})

	t.Run("Lower Middle", func(t *testing.T) {
		cells := NewCellGrid()
		cells[5][4].Alive = true
		count, _, _, _, _ := cells.CountNeighbors(5, 5)
		assert.Equal(t, 1, count)
	})

	t.Run("Lower Right", func(t *testing.T) {
		cells := NewCellGrid()
		cells[6][4].Alive = true
		count, _, _, _, _ := cells.CountNeighbors(5, 5)
		assert.Equal(t, 1, count)
	})
}

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
