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
	req1 := model.Request{
		Floor:     3,
		Direction: model.Up,
	}
	e.AddRequest(req1)
	req2 := model.Request{
		Floor:     7,
		Direction: model.Up,
	}
	e.AddRequest(req2)

	controller.ProcessRequests()

	// Prevent program from exiting immediately
	select {}
}
