package game

import (
	"testing"

	assert "github.com/stretchr/testify/require"
)

func Test_reproduceMajorityWins(t *testing.T) {
	testCases := []struct {
		name   string
		cross  int
		plus   int
		circle int
		dot    int
		want   CellType
	}{
		{"Cross Unanimous", 3, 0, 0, 0, CellType_Cross},
		{"Cross Majority", 2, 1, 0, 0, CellType_Cross},
		{"Cross Minority", 0, 1, 1, 1, CellType_Cross},
		{"Plus Unanimous", 0, 3, 0, 0, CellType_Plus},
		{"Plus Majority", 0, 2, 1, 0, CellType_Plus},
		{"Plus Minority", 1, 0, 1, 1, CellType_Plus},
		{"Circle Unanimous", 0, 0, 3, 0, CellType_Circle},
		{"Circle Majority", 0, 0, 2, 1, CellType_Circle},
		{"Circle Minority", 1, 1, 0, 1, CellType_Circle},
		{"Dot Unanimous", 0, 0, 0, 3, CellType_Dot},
		{"Dot Majority", 1, 0, 0, 2, CellType_Dot},
		{"Dot Minority", 1, 1, 1, 0, CellType_Dot},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cellType := reproduceMajorityWins(tc.cross, tc.plus, tc.circle, tc.dot)
			assert.Equal(t, tc.want, cellType)
		})
	}
}

func Test_reproduceRandomPercentage(t *testing.T) {
	testCases := []struct {
		name   string
		cross  int
		plus   int
		circle int
		dot    int
		want   CellType
	}{
		{"Cross Unanimous", 3, 0, 0, 0, CellType_Cross},
		{"Plus Unanimous", 0, 3, 0, 0, CellType_Plus},
		{"Circle Unanimous", 0, 0, 3, 0, CellType_Circle},
		{"Dot Unanimous", 0, 0, 0, 3, CellType_Dot},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cellType := reproduceRandomPercentage(tc.cross, tc.plus, tc.circle, tc.dot)
			assert.Equal(t, tc.want, cellType)
		})
	}
}

func Benchmark_reproduceMajorityWins_BestCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reproduceMajorityWins(3, 0, 0, 0)
	}
}

func Benchmark_reproduceMajorityWins_WorstCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reproduceMajorityWins(1, 1, 1, 0)
	}
}

func Benchmark_reproduceRandomPercentage_BestCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reproduceRandomPercentage(3, 0, 0, 0)
	}
}

func Benchmark_reproduceRandomPercentage_WorstCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reproduceRandomPercentage(1, 1, 1, 0)
	}
}
