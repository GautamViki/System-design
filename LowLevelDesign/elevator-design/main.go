package main

import (
	"elevator/elevator"
	"elevator/model"
)

func main() {
	// Create a new elevator with 10 floors.
	e := elevator.NewElevator(10)

	// Initialize the controller
	controller := elevator.NewController(e)

	// Add some requests
	e.AddRequest(model.Request{Floor: 3, Direction: model.Up})
	e.AddRequest(model.Request{Floor: 7, Direction: model.Up})
	e.AddRequest(model.Request{Floor: 1, Direction: model.Down})

	controller.ProcessRequests()

	// Prevent program from exiting immediately
	select {}
}
