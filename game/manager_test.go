package game

import (
	"testing"
)

func Benchmark_Manager_swapBuffer(b *testing.B) {
	manager := NewManager()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		manager.swapBuffer()
	}
}

func Benchmark_Manager_updateBuffer_AverageCase(b *testing.B) {
	manager := NewManager()
	cells := manager.Cells()

	for x := range cells {
		if x%2 == 0 {
			for y := range cells[x] {
				cells[x][y].Alive = true
				cells[x][y].Type = CellType(y % 4)
			}
		}
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		manager.updateBuffer()
	}
}

func Benchmark_Manager_updateBuffer_WorstCase(b *testing.B) {
	manager := NewManager()
	cells := manager.Cells()

	for x := range cells {
		for y := range cells[x] {
			cells[x][y].Alive = true
			cells[x][y].Type = CellType(y % 4)
		}
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		manager.updateBuffer()
	}
}
