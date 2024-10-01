package parkinglot

import (
	"car-parking/fee"
	"car-parking/vehicle"
	"errors"
	"time"
)

type ParkingLot struct {
	Spots      []*ParkingSpot
	Ticket     []*ParkingTicket
	NextTicket int
}

func NewParkingLot(smallSpots, mediumSpots, largeSpots int) *ParkingLot {
	spotID := 1
	spots := []*ParkingSpot{}
	for i := 0; i < smallSpots; i++ {
		spots = append(spots, &ParkingSpot{SpotId: spotID, SpotSize: vehicle.Small, IsFree: true})
		spotID++
	}
	for i := 0; i < mediumSpots; i++ {
		spots = append(spots, &ParkingSpot{SpotId: spotID, SpotSize: vehicle.Medium, IsFree: true})
		spotID++
	}
	for i := 0; i < largeSpots; i++ {
		spots = append(spots, &ParkingSpot{SpotId: spotID, SpotSize: vehicle.Large, IsFree: true})
		spotID++
	}
	return &ParkingLot{
		Spots:      spots,
		NextTicket: 1,
	}
}

func (pl *ParkingLot) ParkingVehicle(vehicle *vehicle.Vehicle) (*ParkingTicket, error) {
	spot, err := pl.FindFreeSpot(vehicle)
	if err != nil {
		return nil, err
	}
	err = spot.ParkVehicle(vehicle)
	if err != nil {
		return nil, err
	}

	ticket := &ParkingTicket{
		TicketId:  pl.NextTicket,
		Vehicle:   vehicle,
		Spot:      spot,
		EntryTime: time.Now(),
	}
	pl.Ticket = append(pl.Ticket, ticket)
	pl.NextTicket++
	return ticket, nil
}

func (pl *ParkingLot) RetrieveVehicle(ticketID int, feeCalculator fee.ParkingFeeCalculator) (float64, error) {
	// (Code for vehicle retrieval logic)
	for _, ticket := range pl.Ticket {
		if ticket.TicketId == ticketID && !ticket.IsPaid {
			fee := ticket.CalculateFee(feeCalculator)
			ticket.Spot.RemoveVehicle()
			ticket.IsPaid = true
			return fee, nil
		}
	}
	return 0, errors.New("invalid ticket or already paid")
}

func (pl *ParkingLot) FindFreeSpot(vehicle *vehicle.Vehicle) (*ParkingSpot, error) {
	for _, spot := range pl.Spots {
		if spot.IsFree && spot.SpotSize >= vehicle.VehicleType {
			return spot, nil
		}
	}
	return nil, errors.New("no free spot available")
}
