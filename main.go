package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	rows  = 10
	cols  = 10
	mines = 10
)

type Game struct {
	board         [][]int
	revealedCells [][]bool
	gameOver      bool
	won           bool
}

func NewGame() *Game {
	g := &Game{
		board:         make([][]int, rows),
		revealedCells: make([][]bool, rows),
		gameOver:      false,
		won:           false,
	}
	for i := 0; i < rows; i++ {
		g.board[i] = make([]int, cols)
		g.revealedCells[i] = make([]bool, cols)
	}
	g.placeMines(mines)
	g.calculateNumbers()
	return g
}

func (g *Game) placeMines(numMines int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < numMines; i++ {
		row := rand.Intn(rows)
		col := rand.Intn(cols)
		if g.board[row][col] != -1 {
			g.board[row][col] = -1
		} else {
			i--
		}
	}
}

func (g *Game) calculateNumbers() {
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if g.board[i][j] == -1 {
				continue
			}
			count := 0
			for x := -1; x <= 1; x++ {
				for y := -1; y <= 1; y++ {
					if i+x < 0 || i+x >= rows || j+y < 0 || j+y >= cols {
						continue
					}
					if g.board[i+x][j+y] == -1 {
						count++
					}
				}
			}
			g.board[i][j] = count
		}
	}
}

func (g *Game) reveal(row, col int) {
	if g.board[row][col] == -1 {
		g.gameOver = true
		return
	}
	g.revealedCells[row][col] = true
	if g.board[row][col] == 0 {
		for x := -1; x <= 1; x++ {
			for y := -1; y <= 1; y++ {
				if row+x < 0 || row+x >= rows || col+y < 0 || col+y >= cols {
					continue
				}
				if !g.revealedCells[row+x][col+y] {
					g.reveal(row+x, col+y)
				}
			}
		}
	}
	if g.checkWin() {
		g.won = true
	}
}

func (g *Game) checkWin() bool {
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if !g.revealedCells[i][j] && g.board[i][j] != -1 {
				return false
			}
		}
	}
	return true
}

func (g *Game) PrintBoard() {
	for i := 0; i < len(g.board); i++ {
		for j := 0; j < len(g.board[0]); j++ {
			if g.revealedCells[i][j] {
				if g.board[i][j] == -1 {
					fmt.Print("X ")
				} else if g.board[i][j] == 0 {
					fmt.Print(". ")
				} else {
					fmt.Printf("%d ", g.board[i][j])
				}
			} else {
				fmt.Print("? ")
			}
		}
		fmt.Println()
	}
}

func main() {
	game := NewGame()
	for !game.gameOver && !game.won {
		game.PrintBoard()
		var row, col int
		fmt.Println("Enter row and column to reveal:")
		fmt.Scan(&row, &col)
		game.reveal(row, col)
	}
	game.PrintBoard()
	if game.gameOver {
		fmt.Println("Game over! You hit a mine.")
	} else {
		fmt.Println("Congratulations! You won!")
	}
}
