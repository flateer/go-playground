package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type point struct {
	row int
	col int
}

type TTTBoard [][]string

type AIStrategy func(TTTBoard) point

func main() {
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"-", "-", "-"},
		[]string{"-", "-", "-"},
		[]string{"-", "-", "-"},
	}

	print(board)
	for !gameOver(board) {
		row, col := getMove(board)
		add(board, "X", point{row, col})
		print(board)
		if gameOver(board) {
			fmt.Println("player won")
			return
		}
		getAndMakeAIMove(smartAI, board)
		print(board)
		if gameOver(board) {
			fmt.Println("AI won")
		}

	}
}

func gameOver(board TTTBoard) bool {
	if isFull(board) {
		return true
	}

	for rowIdx, _ := range board {
		if board[rowIdx][0] == board[rowIdx][1] && board[rowIdx][1] == board[rowIdx][2] && board[rowIdx][0] != "-" {
			return true
		}
	}
	for col := 0; col < 3; col++ {
		if board[0][col] == board[1][col] && board[1][col] == board[2][col] && board[0][col] != "-" {
			return true
		}
	}
	if board[0][0] == board[1][1] && board[1][1] == board[2][2] && board[1][1] != "-" {
		return true
	}
	if board[0][2] == board[1][1] && board[1][1] == board[2][0] && board[1][1] != "-" {
		return true
	}
	return false
}

func isFull(board TTTBoard) bool {
	return len(getOpenCoords(board)) == 0
}

func getOpenCoords(board TTTBoard) []point {
	var emptyTiles []point
	for rowIdx, row := range board {
		for colIdx, col := range row {
			if col == "-" {
				emptyTiles = append(emptyTiles, point{rowIdx, colIdx})
			}
		}
	}
	return emptyTiles
}

func getMove(board TTTBoard) (int, int) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter row: ")
	rawRow, _ := reader.ReadString('\n')
	row, err := strconv.Atoi(strings.Trim(rawRow, "\n"))
	if err != nil {
		fmt.Println("Invalid move")
		return getMove(board)
	}

	fmt.Print("Enter column: ")
	rawCol, _ := reader.ReadString('\n')
	col, err := strconv.Atoi(strings.Trim(rawCol, "\n"))
	if err != nil {
		fmt.Println("Invalid move")
		return getMove(board)
	}
	if validateMove(board, row, col) {
		return row, col
	} else {
		fmt.Println("invalid move")
		return getMove(board)
	}
}

func validateMove(board TTTBoard, row int, col int) bool {
	if row < 0 || row > 2 || col < 0 || col > 2 {
		return false
	}
	if board[row][col] != "-" {
		return false
	}
	return true
}

func add(board TTTBoard, token string, move point) {
	board[move.row][move.col] = token
}

func getAndMakeAIMove(strategy AIStrategy, board TTTBoard) {
	move := strategy(board)
	add(board, "0", move)
}

func simpleAI(board TTTBoard) point {
	var empty []point = getOpenCoords(board)
	return empty[0]
}

func smartAI(board TTTBoard) point {
	for _, player := range []string{"O", "X"} {
		for _, point := range getOpenCoords(board) {
			cBoard := copyBoard(board)
			add(cBoard, player, point)
			if gameOver(cBoard) {
				return point
			}
		}
	}
	return simpleAI(board)
}

func copyBoard(board TTTBoard) TTTBoard {
	var newBoard TTTBoard
	for _, row := range board {
		var newRow []string = make([]string, 3)
		copy(newRow, row)
		newBoard = append(newBoard, newRow)
	}
	return newBoard
}

func print(board TTTBoard) {
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], "|"))
	}
	fmt.Println()
}
