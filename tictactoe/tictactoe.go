package tictactoe

import "fmt"

/**
* 3 * 3 matrix
1. A Tic-Tac-Toe board always has exactly 9 cells, so in Go an array is a good choice:
**/

var board [3][3]rune                          // ' ' | 'X' | 'O'
var moves = [9]int{1, 5, 2, 6, 3, 9, 7, 4, 8} // static input: cells 1..9 in play order

func toRC(cell int) (int, int) { return (cell - 1) / 3, (cell - 1) % 3 }

func printBoard() {
	for i := 0; i < 3; i++ {
		fmt.Printf(" %c | %c | %c \n", board[i][0], board[i][1], board[i][2])
		if i < 2 {
			fmt.Println("---+---+---")
		}
	}
	fmt.Println()
}

func winner() rune {
	lines := [8][3][2]int{
		{{0, 0}, {0, 1}, {0, 2}},
		{{1, 0}, {1, 1}, {1, 2}},
		{{2, 0}, {2, 1}, {2, 2}},
		{{0, 0}, {1, 0}, {2, 0}},
		{{0, 1}, {1, 1}, {2, 1}},
		{{0, 2}, {1, 2}, {2, 2}},
		{{0, 0}, {1, 1}, {2, 2}},
		{{0, 2}, {1, 1}, {2, 0}},
	}
	for _, l := range lines {
		a := board[l[0][0]][l[0][1]]
		b := board[l[1][0]][l[1][1]]
		c := board[l[2][0]][l[2][1]]
		if a != ' ' && a == b && b == c {
			return a
		}
	}
	return ' '
}

func Run() {
	// initialize the board with empty spaces
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board[i][j] = ' '
		}
	}

	for count := 0; count < 9; count++ {
		name, mark := "A", 'X'
		if count%2 == 1 {
			name, mark = "B", 'O'
		}

		cell := moves[count] // fetching the move from the static input array
		r, c := toRC(cell)   // convert cell number to row and column

		if cell < 1 || cell > 9 || board[r][c] != ' ' {
			fmt.Printf("User %s: invalid move %d, skipped\n", name, cell)
			continue
		}

		board[r][c] = mark
		fmt.Printf("User %s (%c) -> cell %d\n", name, mark, cell)
		printBoard()

		if w := winner(); w != ' ' {
			fmt.Printf("User %s (%c) wins!\n", name, w)
			return
		}
	}
	fmt.Println("Draw.")
}
