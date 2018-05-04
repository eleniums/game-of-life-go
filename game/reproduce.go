package game

import (
	"math/rand"
)

type ReproduceType int

const (
	ReproduceType_MajorityWins CellType = iota
	ReproduceType_RandomPercentage
)

var (
	ReproduceMethod = ReproduceType_MajorityWins
)

func determineType(cross, plus, circle, dot int) CellType {
	switch ReproduceMethod {
	case ReproduceType_MajorityWins:
		return reproduceMajorityWins(cross, plus, circle, dot)
	case ReproduceType_RandomPercentage:
		return reproduceRandomPercentage(cross, plus, circle, dot)
	default:
		return CellTypeCross
	}
}

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
