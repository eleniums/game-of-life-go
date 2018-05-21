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
		{"Cross Unanimous", 3, 0, 0, 0, CellTypeCross},
		{"Cross Majority", 2, 1, 0, 0, CellTypeCross},
		{"Cross Minority", 0, 1, 1, 1, CellTypeCross},
		{"Plus Unanimous", 0, 3, 0, 0, CellTypePlus},
		{"Plus Majority", 0, 2, 1, 0, CellTypePlus},
		{"Plus Minority", 1, 0, 1, 1, CellTypePlus},
		{"Circle Unanimous", 0, 0, 3, 0, CellTypeCircle},
		{"Circle Majority", 0, 0, 2, 1, CellTypeCircle},
		{"Circle Minority", 1, 1, 0, 1, CellTypeCircle},
		{"Dot Unanimous", 0, 0, 0, 3, CellTypeDot},
		{"Dot Majority", 1, 0, 0, 2, CellTypeDot},
		{"Dot Minority", 1, 1, 1, 0, CellTypeDot},
	}

	for _, tc := range testCases {
		types := []int{tc.cross, tc.plus, tc.circle, tc.dot}
		t.Run(tc.name, func(t *testing.T) {
			cellType := reproduceMajorityWins(types)
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
		{"Cross Unanimous", 3, 0, 0, 0, CellTypeCross},
		{"Plus Unanimous", 0, 3, 0, 0, CellTypePlus},
		{"Circle Unanimous", 0, 0, 3, 0, CellTypeCircle},
		{"Dot Unanimous", 0, 0, 0, 3, CellTypeDot},
	}

	for _, tc := range testCases {
		types := []int{tc.cross, tc.plus, tc.circle, tc.dot}
		t.Run(tc.name, func(t *testing.T) {
			cellType := reproduceRandomPercentage(types)
			assert.Equal(t, tc.want, cellType)
		})
	}
}

func Benchmark_reproduceMajorityWins(b *testing.B) {
	benchmarks := []struct {
		name   string
		cross  int
		plus   int
		circle int
		dot    int
	}{
		{"Best Case", 3, 0, 0, 0},
		{"Worst Case", 1, 1, 1, 0},
	}

	for _, bm := range benchmarks {
		types := []int{bm.cross, bm.plus, bm.circle, bm.dot}
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				reproduceMajorityWins(types)
			}
		})
	}
}

func Benchmark_reproduceRandomPercentage(b *testing.B) {
	benchmarks := []struct {
		name   string
		cross  int
		plus   int
		circle int
		dot    int
	}{
		{"Best Case", 3, 0, 0, 0},
		{"Worst Case", 1, 1, 1, 0},
	}

	for _, bm := range benchmarks {
		types := []int{bm.cross, bm.plus, bm.circle, bm.dot}
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				reproduceRandomPercentage(types)
			}
		})
	}
}
