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
		return CellType_Cross
	}
}

func reproduceMajorityWins(cross, plus, circle, dot int) CellType {
	if cross > 1 {
		return CellType_Cross
	} else if plus > 1 {
		return CellType_Plus
	} else if circle > 1 {
		return CellType_Circle
	} else if dot > 1 {
		return CellType_Dot
	}

	if cross <= 0 {
		return CellType_Cross
	} else if plus <= 0 {
		return CellType_Plus
	} else if circle <= 0 {
		return CellType_Circle
	} else if dot <= 0 {
		return CellType_Dot
	}

	return CellType_Cross
}

func reproduceRandomPercentage(cross, plus, circle, dot int) CellType {
	var types [3]CellType
	index := 0

	for i := 0; i < cross; i++ {
		types[index] = CellType_Cross
		index++
	}
	for i := 0; i < plus; i++ {
		types[index] = CellType_Plus
		index++
	}
	for i := 0; i < circle; i++ {
		types[index] = CellType_Circle
		index++
	}
	for i := 0; i < dot; i++ {
		types[index] = CellType_Dot
		index++
	}

	result := rand.Intn(3)
	return types[result]
}
