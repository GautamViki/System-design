package elevator

import "time"

type Controller struct {
	elevator *Elevator
}

// NewController initializes a new controller.
func NewController(e *Elevator) *Controller {
	return &Controller{
		elevator: e,
	}
}

// ProcessRequests continuously handles elevator requests.
func (c *Controller) ProcessRequests() {
	for {
		c.elevator.Move()
		time.Sleep(1 * time.Second) // Wait between processing steps
	}
}
