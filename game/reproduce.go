package game

func determineType(cross, plus, circle, dot int) CellType {
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
