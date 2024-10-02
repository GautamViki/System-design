package utils

import "fmt"

func DisplayInvalidMove() {
	fmt.Println("Invalid move. Try again.")
}

func DisplayWinner(playerName string) {
	fmt.Printf("Congratulations %s! You are the winner!\n", playerName)
}

func DisplayDraw() {
	fmt.Println("It's a draw!")
}
