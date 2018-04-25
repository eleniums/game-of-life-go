package game

import (
	"testing"
)

func Benchmark_reproduceMajorityWins_BestCase(b *testing.B) {
	typeCounts := []int{3, 0, 0, 0}

	for i := 0; i < b.N; i++ {
		reproduceMajorityWins(typeCounts)
	}
}

func Benchmark_reproduceMajorityWins_WorstCase(b *testing.B) {
	typeCounts := []int{1, 1, 1, 0}

	for i := 0; i < b.N; i++ {
		reproduceMajorityWins(typeCounts)
	}
}

func Benchmark_reproduceRandomPercentage_BestCase(b *testing.B) {
	typeCounts := []int{3, 0, 0, 0}

	for i := 0; i < b.N; i++ {
		reproduceRandomPercentage(typeCounts)
	}
}

func Benchmark_reproduceRandomPercentage_WorstCase(b *testing.B) {
	typeCounts := []int{1, 1, 1, 0}

	for i := 0; i < b.N; i++ {
		reproduceRandomPercentage(typeCounts)
	}
}
