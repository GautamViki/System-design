package game

import (
	"tic-tac-toe/utils"
)

type Game struct {
	Board         *Board
	Player1       *Player
	Player2       *Player
	CurrentPlayer *Player
}

func NewGame(player1Name, player2Name string) *Game {
	board := NewBoard(3)
	player1 := NewPlayer(player1Name, 'X')
	player2 := NewPlayer(player2Name, 'O')

	return &Game{
		Board:         board,
		Player1:       player1,
		Player2:       player2,
		CurrentPlayer: player1,
	}
}

func (g *Game) Start() {
	for !g.Board.IsGameOver() {
		g.Board.PrintBoard()
		g.CurrentPlayer.MakeMove(g.Board)
		if g.Board.HasWinner() {
			utils.DisplayWinner(g.CurrentPlayer.Name)
			return
		}
		g.switchTurn()
	}
	utils.DisplayDraw()
}

func (g *Game) switchTurn() {
	if g.CurrentPlayer == g.Player1 {
		g.CurrentPlayer = g.Player2
	} else {
		g.CurrentPlayer = g.Player1
	}
}
