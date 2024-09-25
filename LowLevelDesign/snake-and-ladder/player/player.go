package player

type Player struct {
	Name     string
	Position int
}

func NewPlayer(name string) *Player {
	return &Player{
		Name:     name,
		Position: 0,
	}
}

func (p *Player) Move(newPosition int) {
	p.Position = newPosition
}
