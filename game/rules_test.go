package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_applyRules_Underpopulation(t *testing.T) {
	alive := applyRules(true, 1)
	assert.False(t, alive)
}

func Test_applyRules_Overpopulation(t *testing.T) {
	alive := applyRules(true, 4)
	assert.False(t, alive)
}

func Test_applyRules_Reproduction(t *testing.T) {
	alive := applyRules(false, 3)
	assert.True(t, alive)
}

func Test_applyRules_StayAlive_2(t *testing.T) {
	alive := applyRules(true, 2)
	assert.True(t, alive)
}

func Test_applyRules_StayAlive_3(t *testing.T) {
	alive := applyRules(true, 3)
	assert.True(t, alive)
}

func Test_applyRules_StayDead_LessThan3(t *testing.T) {
	alive := applyRules(false, 2)
	assert.False(t, alive)
}

func Test_applyRules_StayDead_GreaterThan3(t *testing.T) {
	alive := applyRules(false, 4)
	assert.False(t, alive)
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
