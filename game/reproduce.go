package game

import (
	"math/rand"
)

// ReproduceType is the type of logic used to determine cell type when it becomes alive.
type ReproduceType int

const (
	// ReproduceTypeMajorityWins will determine cell type based on a majority of neighbors.
	ReproduceTypeMajorityWins CellType = iota

	// ReproduceTypeRandomPercentage will determine cell type by randomly picking one of the neighbors.
	ReproduceTypeRandomPercentage
)

var (
	// ReproduceMethod is the method of reproduction used when a cell becomes alive.
	ReproduceMethod = ReproduceTypeMajorityWins
)

// determineType will determine the cell type based on the neighbors, using the ReproduceMethod.
func determineType(types []int) CellType {
	switch ReproduceMethod {
	case ReproduceTypeMajorityWins:
		return reproduceMajorityWins(types)
	case ReproduceTypeRandomPercentage:
		return reproduceRandomPercentage(types)
	default:
		return CellTypeCross
	}
}

// reproduceMajorityWins will determine cell type based on a majority of neighbors.
func reproduceMajorityWins(types []int) CellType {
	for i, v := range types {
		if v > 1 {
			return CellType(i)
		}
	}

	for i, v := range types {
		if v <= 0 {
			return CellType(i)
		}
	}

	return CellTypeCross
}

// reproduceRandomPercentage will determine cell type by randomly picking one of the neighbors.
func reproduceRandomPercentage(types []int) CellType {
	var random [3]CellType
	index := 0

	for i, v := range types {
		for j := 0; j < v; j++ {
			random[index] = CellType(i)
			index++
		}
	}

	result := rand.Intn(3)

	return random[result]
}
