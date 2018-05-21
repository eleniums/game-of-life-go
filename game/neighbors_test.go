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

	manager := NewManager()
	for x := 4; x <= 6; x++ {
		for y := 4; y <= 6; y++ {
			manager.cells.Add(x, y, CellTypeCross)
		}
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			count, _ := manager.countNeighbors(tc.x, tc.y, nil)
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

	manager := NewManager()
	for x := 4; x <= 6; x++ {
		for y := 4; y <= 6; y++ {
			manager.cells.Add(x, y, CellTypeCross)
		}
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, types := manager.countNeighbors(tc.x, tc.y, nil)
			assert.Equal(t, tc.want, types[CellTypeCross])
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

	manager := NewManager()
	for x := 4; x <= 6; x++ {
		for y := 4; y <= 6; y++ {
			manager.cells.Add(x, y, CellTypePlus)
		}
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, types := manager.countNeighbors(tc.x, tc.y, nil)
			assert.Equal(t, tc.want, types[CellTypePlus])
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

	manager := NewManager()
	for x := 4; x <= 6; x++ {
		for y := 4; y <= 6; y++ {
			manager.cells.Add(x, y, CellTypeCircle)
		}
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, types := manager.countNeighbors(tc.x, tc.y, nil)
			assert.Equal(t, tc.want, types[CellTypeCircle])
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

	manager := NewManager()
	for x := 4; x <= 6; x++ {
		for y := 4; y <= 6; y++ {
			manager.cells.Add(x, y, CellTypeDot)
		}
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, types := manager.countNeighbors(tc.x, tc.y, nil)
			assert.Equal(t, tc.want, types[CellTypeDot])
		})
	}
}

func Benchmark_countNeighbors(b *testing.B) {
	b.Run("Best Case", func(b *testing.B) {
		manager := NewManager()

		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			manager.countNeighbors(5, 5, nil)
		}
	})

	b.Run("Worst Case", func(b *testing.B) {
		manager := NewManager()

		manager.cells.Add(4, 5, CellTypeCross)
		manager.cells.Add(6, 5, CellTypeCross)
		manager.cells.Add(5, 6, CellTypeCross)
		manager.cells.Add(4, 6, CellTypeCross)
		manager.cells.Add(6, 6, CellTypeCross)
		manager.cells.Add(5, 4, CellTypeCross)
		manager.cells.Add(4, 4, CellTypeCross)
		manager.cells.Add(6, 4, CellTypeCross)

		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			manager.countNeighbors(5, 5, nil)
		}
	})
}
