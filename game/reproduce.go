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
func determineType(cross, plus, circle, dot int) CellType {
	switch ReproduceMethod {
	case ReproduceTypeMajorityWins:
		return reproduceMajorityWins(cross, plus, circle, dot)
	case ReproduceTypeRandomPercentage:
		return reproduceRandomPercentage(cross, plus, circle, dot)
	default:
		return CellTypeCross
	}
}

// reproduceMajorityWins will determine cell type based on a majority of neighbors.
func reproduceMajorityWins(cross, plus, circle, dot int) CellType {
	if cross > 1 {
		return CellTypeCross
	} else if plus > 1 {
		return CellTypePlus
	} else if circle > 1 {
		return CellTypeCircle
	} else if dot > 1 {
		return CellTypeDot
	}

	if cross <= 0 {
		return CellTypeCross
	} else if plus <= 0 {
		return CellTypePlus
	} else if circle <= 0 {
		return CellTypeCircle
	} else if dot <= 0 {
		return CellTypeDot
	}

	return CellTypeCross
}

// reproduceRandomPercentage will determine cell type by randomly picking one of the neighbors.
func reproduceRandomPercentage(cross, plus, circle, dot int) CellType {
	var types [3]CellType
	index := 0

	for i := 0; i < cross; i++ {
		types[index] = CellTypeCross
		index++
	}
	for i := 0; i < plus; i++ {
		types[index] = CellTypePlus
		index++
	}
	for i := 0; i < circle; i++ {
		types[index] = CellTypeCircle
		index++
	}
	for i := 0; i < dot; i++ {
		types[index] = CellTypeDot
		index++
	}

	result := rand.Intn(3)
	return types[result]
}
