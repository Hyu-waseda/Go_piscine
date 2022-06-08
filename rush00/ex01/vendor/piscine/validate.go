package piscine

func Validate(board []string) bool {
	if IsSquare(board) == false {
		return false
	} else if OneKing(board) == false {
		return false
	} else {
		return true
	}
}

func IsSquare(board []string) bool {
	col := ColLen(board)
	if col == 0 {
		return false
	}
	for _, s := range board {
		if col != RowLen(s) {
			return false
		}
	}
	return true
}

func OneKing(board []string) bool {
	king := false
	for _, s := range board {
		for _, c := range []rune(s) {
			if c == 'K' {
				if king == false {
					king = true
				} else {
					return false
				}
			}
		}
	}
	if king == false {
		return false
	}
	return true
}
