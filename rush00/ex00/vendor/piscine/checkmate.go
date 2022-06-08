package piscine

type Chess struct {
	board map[[2]int]rune
	size  int
	king  [2]int
}

type Move struct {
	dI    []int
	dJ    []int
	ahead bool
	size  int
	piece rune
}

func Checkmate(board []string) bool {
	var chess Chess
	chess.board = convertBoard(board)
	chess.size = strsLen(board)
	chess.king = searchKing(chess.board, chess.size)

	if chess.isInCheck() {
		return true
	} else {
		return false
	}
}

// check if the king is in check
func (chess *Chess) isInCheck() bool {
	for i := 0; i < 4; i++ {
		if chess.checkByPieces(movementOfPieces(i)) {
			return true
		}
	}
	return false
}

func (chess *Chess) checkByPieces(move Move) bool {
	for k := 0; k < move.size; k++ {
		curI, curJ := chess.king[0]+move.dI[k], chess.king[1]+move.dJ[k]
		for isInBoard(curI, curJ, chess.size) {
			if chess.board[[2]int{curI, curJ}] == move.piece {
				return true
			}
			if chess.board[[2]int{curI, curJ}] != '.' || !move.ahead {
				break
			}
			curI, curJ = curI+move.dI[k], curJ+move.dJ[k]
		}
	}
	return false
}

// check if is in board
func isInBoard(i, j, size int) bool {
	if i < 0 || i >= size || j < 0 || j >= size {
		return false
	}
	return true
}

// return the movement of Pawns, Bishops, Rooks, and Queens
func movementOfPieces(i int) Move {
	var move Move
	switch i {
	case 0:
		// Pawns
		move.dI = []int{1, 1}
		move.dJ = []int{-1, 1}
		move.ahead = false
		move.size = 2
		move.piece = 'P'
	case 1:
		// Bishops
		move.dI = []int{-1, 1, 1, -1}
		move.dJ = []int{-1, 1, -1, 1}
		move.ahead = true
		move.size = 4
		move.piece = 'B'
	case 2:
		// Rooks
		move.dI = []int{0, 1, 0, -1}
		move.dJ = []int{1, 0, -1, 0}
		move.ahead = true
		move.size = 4
		move.piece = 'R'
	case 3:
		// Queens
		move.dI = []int{0, 1, 0, -1, 1, -1, 1, -1}
		move.dJ = []int{1, 0, -1, 0, 1, 1, -1, -1}
		move.ahead = true
		move.size = 8
		move.piece = 'Q'
	}
	return move
}

// convert []string to map[[2]int]rune
func convertBoard(board []string) map[[2]int]rune {
	newBoard := map[[2]int]rune{}
	for i, _ := range board {
		for j, c := range []rune(board[i]) {
			if c == '.' || c == 'K' || c == 'P' || c == 'B' || c == 'R' || c == 'Q' {
				newBoard[[2]int{i, j}] = c
			} else {
				newBoard[[2]int{i, j}] = '.'
			}
		}
	}
	return newBoard
}

// measure the length of a []string
func strsLen(strs []string) int {
	var len int
	for range strs {
		len++
	}
	return len
}

// search the king in position the board
func searchKing(board map[[2]int]rune, size int) [2]int {
	var king [2]int
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if board[[2]int{i, j}] == 'K' {
				king[0] = i
				king[1] = j
			}
		}
	}
	return king
}
