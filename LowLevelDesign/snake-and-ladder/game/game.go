package game

import (
	"fmt"
	"sanke-and-ladder/board"
	"sanke-and-ladder/dice"
	"sanke-and-ladder/player"
)

type Game struct {
	Board  *board.Board
	Player []*player.Player
	Dice   *dice.Dice
}

func NewGame(board *board.Board, player []*player.Player, dice *dice.Dice) *Game {
	return &Game{
		Board:  board,
		Player: player,
		Dice:   dice,
	}
}

func (g *Game) Start() {
	for {
		for _, player := range g.Player {
			diceRoll := g.Dice.Roll()
			newPosition := player.Position + diceRoll
			if newPosition > g.Board.Size {
				fmt.Printf("%s rolled %d but can't move beyond board size\n", player.Name, diceRoll)
				continue
			}
			newPosition = g.Board.GetNewPosition(newPosition)
			player.Move(newPosition)
			fmt.Printf("%s rolled a %d and moved to position %d\n", player.Name, diceRoll, newPosition)

			if player.Position == g.Board.Size {
				fmt.Printf("%s wins!\n", player.Name)
				return
			}
		}
	}
}
