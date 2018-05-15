package game

import (
	"testing"

	assert "github.com/stretchr/testify/require"
)

func Test_countNeighbors_Alive(t *testing.T) {
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

	cells := NewGrid()
	for x := 4; x <= 6; x++ {
		for y := 4; y <= 6; y++ {
			cells.Add(x, y, CellTypeCross)
		}
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			count, _, _, _, _ := countNeighbors(cells, tc.x, tc.y)
			assert.Equal(t, tc.want, count)
		})
	}
}

func Test_countNeighbors_Type_Cross(t *testing.T) {
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

	cells := NewGrid()
	for x := 4; x <= 6; x++ {
		for y := 4; y <= 6; y++ {
			cells.Add(x, y, CellTypeCross)
		}
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, cross, _, _, _ := countNeighbors(cells, tc.x, tc.y)
			assert.Equal(t, tc.want, cross)
		})
	}
}

func Test_countNeighbors_Type_Plus(t *testing.T) {
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

	cells := NewGrid()
	for x := 4; x <= 6; x++ {
		for y := 4; y <= 6; y++ {
			cells.Add(x, y, CellTypePlus)
		}
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, _, plus, _, _ := countNeighbors(cells, tc.x, tc.y)
			assert.Equal(t, tc.want, plus)
		})
	}
}

func Test_countNeighbors_Type_Circle(t *testing.T) {
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

	cells := NewGrid()
	for x := 4; x <= 6; x++ {
		for y := 4; y <= 6; y++ {
			cells.Add(x, y, CellTypeCircle)
		}
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, _, _, circle, _ := countNeighbors(cells, tc.x, tc.y)
			assert.Equal(t, tc.want, circle)
		})
	}
}

func Test_countNeighbors_Type_Dot(t *testing.T) {
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

	cells := NewGrid()
	for x := 4; x <= 6; x++ {
		for y := 4; y <= 6; y++ {
			cells.Add(x, y, CellTypeDot)
		}
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, _, _, _, dot := countNeighbors(cells, tc.x, tc.y)
			assert.Equal(t, tc.want, dot)
		})
	}
}

func Benchmark_countNeighbors(b *testing.B) {
	b.Run("Best Case", func(b *testing.B) {
		cells := NewGrid()

		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			countNeighbors(cells, 5, 5)
		}
	})

	b.Run("Worst Case", func(b *testing.B) {
		cells := NewGrid()

		cells.Add(4, 5, CellTypeCross)
		cells.Add(6, 5, CellTypeCross)
		cells.Add(5, 6, CellTypeCross)
		cells.Add(4, 6, CellTypeCross)
		cells.Add(6, 6, CellTypeCross)
		cells.Add(5, 4, CellTypeCross)
		cells.Add(4, 4, CellTypeCross)
		cells.Add(6, 4, CellTypeCross)

		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			countNeighbors(cells, 5, 5)
		}
	})
}
