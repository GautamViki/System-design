package model

// Direction represents the direction the elevator is moving.
type Direction int

const (
	Idle Direction = iota
	Up
	Down
)

// Request represents an elevator request.
type Request struct {
	Floor     int
	Direction Direction
}
