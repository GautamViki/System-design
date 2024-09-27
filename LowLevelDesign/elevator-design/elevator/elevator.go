package elevator

import (
	"elevator/model"
	"fmt"
	"time"
)

type Elevator struct {
	currentFloor int
	direction    model.Direction
	request      []model.Request
	numFloor     int
}

func NewElevator(numFloor int) *Elevator {
	return &Elevator{
		currentFloor: 0,
		direction:    model.Idle,
		numFloor:     numFloor,
		request:      []model.Request{},
	}
}

func (e *Elevator) AddRequest(req model.Request) {
	e.request = append(e.request, req)
}

// Move moves the elevator based on the requests.
func (e *Elevator) Move() {
	if len(e.request) == 0 {
		fmt.Println("No requests, elevator is idle.")
		return
	}

	// Move towards the first request
	target := e.request[0].Floor
	movement := ""
	if e.currentFloor < target {
		e.direction = model.Up
		movement = "up"
		e.currentFloor++
	} else if e.currentFloor > target {
		e.direction = model.Down
		movement = "down"
		e.currentFloor--
	}

	fmt.Printf("Elevator is at floor %d and going %s\n", e.currentFloor, movement)

	// Simulate some time to move between floors
	time.Sleep(1 * time.Second)

	// If at the target floor, open doors and remove request
	if e.currentFloor == target {
		fmt.Println("Elevator has arrived. Opening doors.")
		time.Sleep(2 * time.Second)
		e.request = e.request[1:]
		fmt.Println("Doors closing.")
	}
}
