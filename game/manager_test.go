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

func Benchmark_Manager_Update(b *testing.B) {
	b.Run("Zero Cells", func(b *testing.B) {
		manager := NewManager()

		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			manager.updateBuffer()
			manager.swapBuffer()
		}
	})

	b.Run("Infinite 1", func(b *testing.B) {
		manager := NewManager()

		manager.cells.Add(5, 5, CellTypeCross)
		manager.cells.Add(7, 5, CellTypeCross)
		manager.cells.Add(7, 6, CellTypeCross)
		manager.cells.Add(9, 7, CellTypeCross)
		manager.cells.Add(9, 8, CellTypeCross)
		manager.cells.Add(9, 9, CellTypeCross)
		manager.cells.Add(11, 8, CellTypeCross)
		manager.cells.Add(11, 9, CellTypeCross)
		manager.cells.Add(11, 10, CellTypeCross)
		manager.cells.Add(12, 9, CellTypeCross)

		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			manager.updateBuffer()
			manager.swapBuffer()
		}
	})

	b.Run("Infinite 2", func(b *testing.B) {
		manager := NewManager()

		manager.cells.Add(5, 5, CellTypeCross)
		manager.cells.Add(7, 5, CellTypeCross)
		manager.cells.Add(9, 5, CellTypeCross)
		manager.cells.Add(6, 6, CellTypeCross)
		manager.cells.Add(7, 6, CellTypeCross)
		manager.cells.Add(9, 6, CellTypeCross)
		manager.cells.Add(8, 7, CellTypeCross)
		manager.cells.Add(9, 7, CellTypeCross)
		manager.cells.Add(5, 8, CellTypeCross)
		manager.cells.Add(5, 9, CellTypeCross)
		manager.cells.Add(6, 9, CellTypeCross)
		manager.cells.Add(7, 9, CellTypeCross)
		manager.cells.Add(9, 9, CellTypeCross)

		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			manager.updateBuffer()
			manager.swapBuffer()
		}
	})

	b.Run("Glider Gun", func(b *testing.B) {
		manager := NewManager()

		manager.cells.Add(5, 8, CellTypeCross)
		manager.cells.Add(5, 9, CellTypeCross)
		manager.cells.Add(6, 8, CellTypeCross)
		manager.cells.Add(6, 9, CellTypeCross)

		manager.cells.Add(15, 7, CellTypeCross)
		manager.cells.Add(15, 8, CellTypeCross)
		manager.cells.Add(15, 9, CellTypeCross)
		manager.cells.Add(16, 6, CellTypeCross)
		manager.cells.Add(16, 10, CellTypeCross)
		manager.cells.Add(17, 5, CellTypeCross)
		manager.cells.Add(17, 11, CellTypeCross)
		manager.cells.Add(18, 5, CellTypeCross)
		manager.cells.Add(18, 11, CellTypeCross)
		manager.cells.Add(19, 8, CellTypeCross)
		manager.cells.Add(20, 6, CellTypeCross)
		manager.cells.Add(20, 10, CellTypeCross)
		manager.cells.Add(21, 7, CellTypeCross)
		manager.cells.Add(21, 8, CellTypeCross)
		manager.cells.Add(21, 9, CellTypeCross)
		manager.cells.Add(22, 8, CellTypeCross)

		manager.cells.Add(25, 9, CellTypeCross)
		manager.cells.Add(25, 10, CellTypeCross)
		manager.cells.Add(25, 11, CellTypeCross)
		manager.cells.Add(26, 9, CellTypeCross)
		manager.cells.Add(26, 10, CellTypeCross)
		manager.cells.Add(26, 11, CellTypeCross)
		manager.cells.Add(27, 8, CellTypeCross)
		manager.cells.Add(27, 12, CellTypeCross)
		manager.cells.Add(29, 7, CellTypeCross)
		manager.cells.Add(29, 8, CellTypeCross)
		manager.cells.Add(29, 12, CellTypeCross)
		manager.cells.Add(29, 13, CellTypeCross)

		manager.cells.Add(39, 10, CellTypeCross)
		manager.cells.Add(39, 11, CellTypeCross)
		manager.cells.Add(40, 10, CellTypeCross)
		manager.cells.Add(40, 11, CellTypeCross)

		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			manager.updateBuffer()
			manager.swapBuffer()
		}
	})
}
