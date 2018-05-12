package game

import (
	"testing"

	assert "github.com/stretchr/testify/require"
)

func Test_Manager_swapBuffer(t *testing.T) {
	manager := NewManager()

	manager.cells[25][25].Alive = true
	manager.buffer[5][5].Alive = true

	manager.swapBuffer()

	assert.False(t, manager.cells[25][25].Alive)
	assert.True(t, manager.cells[5][5].Alive)

	assert.True(t, manager.buffer[25][25].Alive)
	assert.False(t, manager.buffer[5][5].Alive)
}

func Test_Manager_updateBuffer(t *testing.T) {
	manager := NewManager()
	cells := manager.Cells()

	cells[24][25].Alive = true
	cells[25][25].Alive = true
	cells[26][25].Alive = true

	manager.updateBuffer()

	assert.False(t, manager.buffer[24][25].Alive)
	assert.True(t, manager.buffer[25][25].Alive)
	assert.False(t, manager.buffer[26][25].Alive)
	assert.True(t, manager.buffer[25][24].Alive)
	assert.True(t, manager.buffer[25][26].Alive)
}

func Benchmark_Manager_swapBuffer(b *testing.B) {
	manager := NewManager()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		manager.swapBuffer()
	}
}

func Benchmark_Manager_updateBuffer(b *testing.B) {
	b.Run("Best Case", func(b *testing.B) {
		manager := NewManager()

		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			manager.updateBuffer()
		}
	})

	b.Run("Average Case", func(b *testing.B) {
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
	})

	b.Run("Worst Case", func(b *testing.B) {
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
	})
}
