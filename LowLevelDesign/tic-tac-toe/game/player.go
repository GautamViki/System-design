package game

import "tic-tac-toe/utils"

type Player struct {
	Name string
	Mark rune
}

func NewPlayer(name string, mark rune) *Player {
	return &Player{
		Name: name,
		Mark: mark,
	}
}

func (p *Player) MakeMove(board *Board) {
	for {
		x, y := utils.ReadPlayerInput(p.Name)
		if board.IsMoveValid(x, y) {
			board.PlaceMark(x, y, p.Mark)
			return
		}
		utils.DisplayInvalidMove()
	}
}
