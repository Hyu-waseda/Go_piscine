package piscine

import "ft"

func EightQueens() {
	var board [8][8]bool
	initBoard(&board)
	solve(0, board)
}

func initBoard(board *[8][8]bool) {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			(*board)[i][j] = false
		}
	}
}

func solve(j int, board [8][8]bool) {
	if j == 8 {
		printQueens(board)
		return
	}
	for i := 0; i < 8; i++ {
		if canPutQueen(i, j, board) {
			board[i][j] = true
			solve(j + 1, board)
			board[i][j] = false
		}
	}
}

func canPutQueen(i, j int, board [8][8]bool) (bool){
	if checkLeft(i, j, board) &&
		checkUpperLeft(i, j, board) &&
		checkLowerLeft(i, j, board) {
			return true
	}
	return false
}

func checkLeft(i, j int, board [8][8]bool) (bool){
	for k := 0; k < j; k++ {
		if board[i][k] == true {
			return false
		}
	}
	return true
}

func checkUpperLeft(i, j int, board [8][8]bool) (bool){
	for k, l := i, j; 0 <= k && 0 <= l; k, l = k - 1, l - 1 {
		if board[k][l] == true {
			return false
		}
	}
	return true
}

func checkLowerLeft(i, j int, board [8][8]bool) (bool){
	for k, l := i, j; k < 8 && 0 <= l; k, l = k + 1, l - 1 {
		if board[k][l] == true {
			return false
		}
	}
	return true
}

func printQueens(board [8][8]bool) {
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
