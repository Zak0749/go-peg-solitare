package main

import "fmt"

type Tile struct {
	exists bool
	marble bool
}

type Position struct {
	x int
	y int
}

type Board [][]Tile

func pretyify(board Board) string {
	str := ""
	for _, col := range board {
		for _, val := range col {
			if !val.exists {
				str += " "
			} else if val.marble {
				str += "•"
			} else {
				str += "o"
			}
		}

		str += "\n"
	}

	return str
}

func existsAnd(condition bool, board Board, x int, y int) bool {
	return board[y][x].exists && board[x][y].marble == condition
}

func edit(board Board, old Position, jump Position, new Position) Board {
	newBoard := board
	newBoard[old.y][old.x].marble = false
	newBoard[jump.y][jump.x].marble = false
	newBoard[new.y][new.x].marble = true
	return newBoard
}

func getMoves(board Board) []Board {
	moves := []Board{}
	for x, col := range board {
		for y, val := range col {
			if !val.exists || !val.marble {
				continue
			}

			if y > 1 && existsAnd(true, board, x, y-1) && existsAnd(false, board, x, y-2) {
				moves = append(moves, edit(board, Position{x: x, y: y}, Position{x: x, y: y - 1}, Position{x: x, y: y - 2}))
			}

			if y < 7 && existsAnd(true, board, x, y+1) && existsAnd(false, board, x, y+2) {
				moves = append(moves, edit(board, Position{x: x, y: y}, Position{x: x, y: y + 1}, Position{x: x, y: y + 2}))
			}

			if x > 1 && existsAnd(true, board, x-1, y) && existsAnd(false, board, x-2, y) {
				moves = append(moves, edit(board, Position{x: x, y: y}, Position{x: x - 1, y: y}, Position{x: x - 2, y: y}))
			}

			if x < 7 && existsAnd(true, board, x+1, y) && existsAnd(false, board, x+2, y) {
				moves = append(moves, edit(board, Position{x: x, y: y}, Position{x: x + 1, y: y}, Position{x: x + 2, y: y}))
			}
		}
	}

	return moves
}

func main() {

	board := Board{
		{
			{exists: false, marble: false},
			{exists: false, marble: false},
			{exists: false, marble: false},
			{exists: true, marble: true},
			{exists: true, marble: true},
			{exists: true, marble: true},
			{exists: false, marble: false},
			{exists: false, marble: false},
			{exists: false, marble: false},
		},
		{
			{exists: false, marble: false},
			{exists: false, marble: false},
			{exists: false, marble: false},
			{exists: true, marble: true},
			{exists: true, marble: true},
			{exists: true, marble: true},
			{exists: false, marble: false},
			{exists: false, marble: false},
			{exists: false, marble: false},
		},
		{
			{exists: false, marble: false},
			{exists: false, marble: false},
			{exists: false, marble: false},
			{exists: true, marble: true},
			{exists: true, marble: true},
			{exists: true, marble: true},
			{exists: false, marble: false},
			{exists: false, marble: false},
			{exists: false, marble: false},
		},
		{
			{exists: true, marble: true},
			{exists: true, marble: true},
			{exists: true, marble: true},
			{exists: true, marble: true},
			{exists: true, marble: true},
			{exists: true, marble: true},
			{exists: true, marble: true},
			{exists: true, marble: true},
			{exists: true, marble: true},
		},
		{
			{exists: true, marble: true},
			{exists: true, marble: true},
			{exists: true, marble: true},
			{exists: true, marble: true},
			{exists: true, marble: false},
			{exists: true, marble: true},
			{exists: true, marble: true},
			{exists: true, marble: true},
			{exists: true, marble: true},
		},
		{
			{exists: true, marble: true},
			{exists: true, marble: true},
			{exists: true, marble: true},
			{exists: true, marble: true},
			{exists: true, marble: true},
			{exists: true, marble: true},
			{exists: true, marble: true},
			{exists: true, marble: true},
			{exists: true, marble: true},
		},
		{
			{exists: false, marble: false},
			{exists: false, marble: false},
			{exists: false, marble: false},
			{exists: true, marble: true},
			{exists: true, marble: true},
			{exists: true, marble: true},
			{exists: false, marble: false},
			{exists: false, marble: false},
			{exists: false, marble: false},
		},
		{
			{exists: false, marble: false},
			{exists: false, marble: false},
			{exists: false, marble: false},
			{exists: true, marble: true},
			{exists: true, marble: true},
			{exists: true, marble: true},
			{exists: false, marble: false},
			{exists: false, marble: false},
			{exists: false, marble: false},
		},
		{
			{exists: false, marble: false},
			{exists: false, marble: false},
			{exists: false, marble: false},
			{exists: true, marble: true},
			{exists: true, marble: true},
			{exists: true, marble: true},
			{exists: false, marble: false},
			{exists: false, marble: false},
			{exists: false, marble: false},
		},
	}

	moves := getMoves(board)

	for _, val := range moves {
		fmt.Println(pretyify(val))
	}

}
