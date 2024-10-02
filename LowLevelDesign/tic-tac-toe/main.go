package main

import (
	"fmt"
	"tic-tac-toe/game"
)

func main() {
	fmt.Println("Welcome to Tic Tac Toe!")

	// Initialize a new game with two players
	gameInstance := game.NewGame("Player 1", "Player 2")

	// Start the game
	gameInstance.Start()
}
