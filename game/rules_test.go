package game

import (
	"testing"

	assert "github.com/stretchr/testify/require"
)

func Test_applyRules(t *testing.T) {
	testCases := []struct {
		name      string
		alive     bool
		neighbors int
		want      bool
	}{
		{"Underpopulation", true, 1, false},
		{"Overpopulation", true, 4, false},
		{"Reproduction", false, 3, true},
		{"Stay Alive With 2", true, 2, true},
		{"Stay Alive With 3", true, 3, true},
		{"Stay Dead With Less Than 3", false, 2, false},
		{"Stay Dead With Greater Than 3", false, 4, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			alive := applyRules(tc.alive, tc.neighbors)
			assert.Equal(t, tc.want, alive)
		})
	}
}

func Benchmark_applyRules(b *testing.B) {
	benchmarks := []struct {
		name      string
		alive     bool
		neighbors int
	}{
		{"Underpopulation", true, 1},
		{"Overpopulation", true, 4},
		{"Reproduction", false, 3},
		{"Stay Alive With 2", true, 2},
		{"Stay Alive With 3", true, 3},
		{"Stay Dead With Less Than 3", false, 2},
		{"Stay Dead With Greater Than 3", false, 4},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				applyRules(bm.alive, bm.neighbors)
			}
		})
	}
}
