package game

import (
	"testing"

	assert "github.com/stretchr/testify/require"
)

func Test_Manager_swapBuffer(t *testing.T) {
	manager := NewManager()

	manager.cells.Add(25, 25, CellTypeCross)
	manager.buffer.Add(5, 5, CellTypeCross)

	manager.swapBuffer()

	var ok bool
	_, ok = manager.cells.Retrieve(25, 25)
	assert.False(t, ok)
	_, ok = manager.cells.Retrieve(5, 5)
	assert.True(t, ok)

	_, ok = manager.buffer.Retrieve(25, 25)
	assert.True(t, ok)
	_, ok = manager.buffer.Retrieve(5, 5)
	assert.False(t, ok)
}

func Test_Manager_updateBuffer(t *testing.T) {
	manager := NewManager()

	manager.cells.Add(24, 25, CellTypeCross)
	manager.cells.Add(25, 25, CellTypeCross)
	manager.cells.Add(26, 25, CellTypeCross)

	manager.updateBuffer()

	var ok bool
	_, ok = manager.buffer.Retrieve(24, 25)
	assert.False(t, ok)
	_, ok = manager.buffer.Retrieve(25, 25)
	assert.True(t, ok)
	_, ok = manager.buffer.Retrieve(26, 25)
	assert.False(t, ok)
	_, ok = manager.buffer.Retrieve(25, 24)
	assert.True(t, ok)
	_, ok = manager.buffer.Retrieve(25, 26)
	assert.True(t, ok)
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

		for x := 0; x < 288; x++ {
			if x%2 == 0 {
				for y := 0; y < 288; y++ {
					manager.cells.Add(x, y, CellType(y%4))
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

		for x := 0; x < 288; x++ {
			for y := 0; y < 288; y++ {
				manager.cells.Add(x, y, CellType(y%4))
			}
		}

		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			manager.updateBuffer()
		}
	})
}
