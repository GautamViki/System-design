package utils

import "fmt"

func ReadPlayerInput(playerName string) (int, int) {
    var x, y int
    fmt.Printf("%s, enter your move (row and column): ", playerName)
    fmt.Scan(&x, &y)
    return x, y
}