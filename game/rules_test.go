package game

import (
	"testing"
)

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
