package piscine

import "ft"

func EightQueens() {
	var board [8][8]bool
	InitBoard(&board)
	Solve(0, board)
}

func InitBoard(board *[8][8]bool) {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			(*board)[i][j] = false
		}
	}
}

func Solve(j int, board [8][8]bool) {
	if j == 8 {
		PrintQueens(board)
		return
	}
	for i := 0; i < 8; i++ {
		if CanPutQueen(i, j, board) {
			board[i][j] = true
			Solve(j + 1, board)
			board[i][j] = false
		}
	}
}

func CanPutQueen(i int, j int, board [8][8]bool) (bool){
	for k := 0; k < j; k++ {
		if board[i][k] == true {
			return false
		}
	}
	for k, l := i, j; 0 <= k && 0 <= l; k, l = k - 1, l - 1 {
		if board[k][l] == true {
			return false
		}
	}
	for k, l := i, j; k < 8 && 0 <= l; k, l = k + 1, l - 1 {
		if board[k][l] == true {
			return false
		}
	}
	return true
}

func PrintQueens(board [8][8]bool) {
	for j := 0; j < 8; j++ {
		for i := 0; i < 8; i++ {
			if board[i][j] == true {
				ft.PrintRune(rune(int('1') + i))
				break
			}
		}
	}
	ft.PrintRune('\n')
}
