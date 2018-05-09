package game

import (
	"testing"

	assert "github.com/stretchr/testify/require"
)

func Test_CellGrid_CountNeighbors_Alive(t *testing.T) {
	testCases := []struct {
		name string
		x    int
		y    int
		want int
	}{
		{"Upper Left", 4, 6, 3},
		{"Upper Middle", 5, 6, 5},
		{"Upper Right", 6, 6, 3},
		{"Middle Left", 4, 5, 5},
		{"Center", 5, 5, 8},
		{"Middle Right", 6, 5, 5},
		{"Lower Left", 4, 4, 3},
		{"Lower Middle", 5, 4, 5},
		{"Lower Right", 6, 4, 3},
	}

	cells := NewCellGrid()
	for x := 4; x <= 6; x++ {
		for y := 4; y <= 6; y++ {
			cells[x][y].Alive = true
		}
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			count, _, _, _, _ := cells.CountNeighbors(tc.x, tc.y)
			assert.Equal(t, tc.want, count)
		})
	}
}

func Test_CellGrid_CountNeighbors_Type_Cross(t *testing.T) {
	testCases := []struct {
		name string
		x    int
		y    int
		want int
	}{
		{"Upper Left", 4, 6, 3},
		{"Upper Middle", 5, 6, 5},
		{"Upper Right", 6, 6, 3},
		{"Middle Left", 4, 5, 5},
		{"Center", 5, 5, 8},
		{"Middle Right", 6, 5, 5},
		{"Lower Left", 4, 4, 3},
		{"Lower Middle", 5, 4, 5},
		{"Lower Right", 6, 4, 3},
	}

	cells := NewCellGrid()
	for x := 4; x <= 6; x++ {
		for y := 4; y <= 6; y++ {
			cells[x][y].Alive = true
			cells[x][y].Type = CellTypeCross
		}
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, cross, _, _, _ := cells.CountNeighbors(tc.x, tc.y)
			assert.Equal(t, tc.want, cross)
		})
	}
}

func Test_CellGrid_CountNeighbors_Type_Plus(t *testing.T) {
	testCases := []struct {
		name string
		x    int
		y    int
		want int
	}{
		{"Upper Left", 4, 6, 3},
		{"Upper Middle", 5, 6, 5},
		{"Upper Right", 6, 6, 3},
		{"Middle Left", 4, 5, 5},
		{"Center", 5, 5, 8},
		{"Middle Right", 6, 5, 5},
		{"Lower Left", 4, 4, 3},
		{"Lower Middle", 5, 4, 5},
		{"Lower Right", 6, 4, 3},
	}

	cells := NewCellGrid()
	for x := 4; x <= 6; x++ {
		for y := 4; y <= 6; y++ {
			cells[x][y].Alive = true
			cells[x][y].Type = CellTypePlus
		}
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, _, plus, _, _ := cells.CountNeighbors(tc.x, tc.y)
			assert.Equal(t, tc.want, plus)
		})
	}
}

func Test_CellGrid_CountNeighbors_Type_Circle(t *testing.T) {
	testCases := []struct {
		name string
		x    int
		y    int
		want int
	}{
		{"Upper Left", 4, 6, 3},
		{"Upper Middle", 5, 6, 5},
		{"Upper Right", 6, 6, 3},
		{"Middle Left", 4, 5, 5},
		{"Center", 5, 5, 8},
		{"Middle Right", 6, 5, 5},
		{"Lower Left", 4, 4, 3},
		{"Lower Middle", 5, 4, 5},
		{"Lower Right", 6, 4, 3},
	}

	cells := NewCellGrid()
	for x := 4; x <= 6; x++ {
		for y := 4; y <= 6; y++ {
			cells[x][y].Alive = true
			cells[x][y].Type = CellTypeCircle
		}
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, _, _, circle, _ := cells.CountNeighbors(tc.x, tc.y)
			assert.Equal(t, tc.want, circle)
		})
	}
}

func Test_CellGrid_CountNeighbors_Type_Dot(t *testing.T) {
	testCases := []struct {
		name string
		x    int
		y    int
		want int
	}{
		{"Upper Left", 4, 6, 3},
		{"Upper Middle", 5, 6, 5},
		{"Upper Right", 6, 6, 3},
		{"Middle Left", 4, 5, 5},
		{"Center", 5, 5, 8},
		{"Middle Right", 6, 5, 5},
		{"Lower Left", 4, 4, 3},
		{"Lower Middle", 5, 4, 5},
		{"Lower Right", 6, 4, 3},
	}

	cells := NewCellGrid()
	for x := 4; x <= 6; x++ {
		for y := 4; y <= 6; y++ {
			cells[x][y].Alive = true
			cells[x][y].Type = CellTypeDot
		}
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, _, _, _, dot := cells.CountNeighbors(tc.x, tc.y)
			assert.Equal(t, tc.want, dot)
		})
	}
}

func Benchmark_CellGrid_CountNeighbors(b *testing.B) {
	b.Run("Best Case", func(b *testing.B) {
		cells := NewCellGrid()

		cells[4][5].Alive = false
		cells[6][5].Alive = false
		cells[5][6].Alive = false
		cells[4][6].Alive = false
		cells[6][6].Alive = false
		cells[5][4].Alive = false
		cells[4][4].Alive = false
		cells[6][4].Alive = false

		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			cells.CountNeighbors(5, 5)
		}
	})

	b.Run("Worst Case", func(b *testing.B) {
		cells := NewCellGrid()

		cells[4][5].Alive = true
		cells[6][5].Alive = true
		cells[5][6].Alive = true
		cells[4][6].Alive = true
		cells[6][6].Alive = true
		cells[5][4].Alive = true
		cells[4][4].Alive = true
		cells[6][4].Alive = true

		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			cells.CountNeighbors(5, 5)
		}
	})
}
