package main

import (
	"sanke-and-ladder/board"
	"sanke-and-ladder/dice"
	"sanke-and-ladder/game"
	"sanke-and-ladder/player"
)

func main() {
	snakes := []*board.Snake{
		board.NewSnake(17, 7),
		board.NewSnake(54, 34),
		board.NewSnake(62, 19),
	}

	ladders := []*board.Ladder{
		board.NewLadder(3, 22),
		board.NewLadder(5, 8),
		board.NewLadder(20, 38),
	}

	// Create the board, dice, and players
	gameBoard := board.NewBoard(100, snakes, ladders)
	gameDice := dice.NewDice(6)
	players := []*player.Player{
		player.NewPlayer("Alice"),
		player.NewPlayer("Bob"),
	}

	// Start the game
	gameInstance := game.NewGame(gameBoard, players, gameDice)
	gameInstance.Start()
}
