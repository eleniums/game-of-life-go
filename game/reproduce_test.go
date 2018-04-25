package game

import (
	"testing"
)

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
