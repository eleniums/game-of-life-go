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

func Benchmark_applyRules_Underpopulation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		applyRules(true, 1)
	}
}

func Benchmark_applyRules_Overpopulation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		applyRules(true, 4)
	}
}

func Benchmark_applyRules_Reproduction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		applyRules(false, 3)
	}
}

func Benchmark_applyRules_StayAlive_2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		applyRules(true, 2)
	}
}

func Benchmark_applyRules_StayAlive_3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		applyRules(true, 3)
	}
}

func Benchmark_applyRules_StayDead_LessThan3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		applyRules(false, 2)
	}
}

func Benchmark_applyRules_StayDead_GreaterThan3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		applyRules(false, 4)
	}
}
