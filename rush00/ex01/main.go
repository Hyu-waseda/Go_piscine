package main

import (
	"bufio"
	"fmt"
	"ft"
	"os"
	"piscine"
	"math/rand"
	"time"
)

func main() {
	board := []string{
		"R...",
		".K..",
		"..P.",
		"....",
	}
	for i := 0; i < 10; i++ {
		rand.Seed(time.Now().UnixNano())
		randomChange(&board)
		//printBoard(board)
		Judge(board)
		fmt.Println("-----------")
	}
	//Judge(board)
}

func randomChange(board *[]string) {
	var newboard []string
	for _, s := range *board {
		new := ""
		for i, v := range s {
			i++
			n := rand.Intn(20)
			if v == 'K' {
				new += "K"
			} else if n == 1 {
				new += "Q"
			} else if n == 2 {
				new += "B"
			} else if n == 3 {
				new += "R"
			} else if n == 4 {
				new += "P"
			} else {
				new += "."
			}
		}
		newboard = append(newboard, new)
	}
	*board = newboard
}

func Judge(board []string) {
	board = MakeBoard(board)
	printBoard(board)
	if !piscine.Validate(board) {
		printlnStr("Input error")
		return
	}
	result := piscine.Judge(board)
	printlnStr(result)
}

func MakeBoard(board []string) []string {
	var newBoard []string
	var createNew bool
	for i, s := range os.Args {
		if i == 0 {
			continue
		}
		if i == 1 && s == "-c" {
			createNew = true
			continue
		} else if i == 1 && s == "-i" {
			return ReadBoard()
		}
		newBoard = append(newBoard, s)
	}
	if createNew {
		return newBoard
	} else {
		return board
	}
}

func ReadBoard() []string {
	var board []string
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		s := scanner.Text()
		if s == "" {
			break
		}
		board = append(board, s)
	}
	return board
}

func printBoard(board []string) {
	for _, s := range board {
		fmt.Println(s)
	}
}

func printlnStr(str string) {
	for _, c := range str {
		ft.PrintRune(c)
	}
	ft.PrintRune('\n')
}
