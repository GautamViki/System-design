package board

type Snake struct {
	Head int
	Tail int
}

func NewSnake(head, tail int) *Snake {
	return &Snake{
		Head: head,
		Tail: tail}
}
