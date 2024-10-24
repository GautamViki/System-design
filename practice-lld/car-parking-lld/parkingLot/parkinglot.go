package parkinglot

import (
	feecalculator "car-parking/feeCalculator"
	parkingslot "car-parking/parkingSlot"
	parkingticket "car-parking/parkingTicket"
	"car-parking/vehicle"
	"errors"
	"time"
)

type Lot struct {
	Spots      []*parkingslot.Spot
	Tickets    []*parkingticket.Ticket
	NextTicket int
}

func NewParkingLot(smallSpots, mediumSpots, largeSpots int) *Lot {
	spots := []*parkingslot.Spot{}
	count := 1
	for count <= smallSpots {
		spots = append(spots, &parkingslot.Spot{
			SpotId: count,
			Size:   vehicle.Small,
			IsFree: true,
		})
		count++
	}
	i := 0
	for i < mediumSpots {
		spots = append(spots, &parkingslot.Spot{
			SpotId: count,
			Size:   vehicle.Medium,
			IsFree: true,
		})
		i++
		count++
	}
	i = 0
	for i < mediumSpots {
		spots = append(spots, &parkingslot.Spot{
			SpotId: count,
			Size:   vehicle.Large,
			IsFree: true,
		})
		i++
		count++
	}
	return &Lot{
		Spots:      spots,
		NextTicket: 1,
	}
}

func (l *Lot) ParkedVehicle(vehicle *vehicle.Vehicle) (*parkingticket.Ticket, error) {
	spot, err := l.IsFreeSpot(vehicle)
	if err != nil {
		return &parkingticket.Ticket{}, err
	}
	err = spot.ParkCar(vehicle)
	if err != nil {
		return &parkingticket.Ticket{}, err
	}
	ticket := &parkingticket.Ticket{
		Id:        l.NextTicket,
		Spot:      spot,
		EntryTime: time.Now(),
		Vehicle:   *vehicle,
	}
	l.Tickets = append(l.Tickets, ticket)
	l.NextTicket++
	return ticket, nil
}

func (l *Lot) IsFreeSpot(vehicle *vehicle.Vehicle) (*parkingslot.Spot, error) {
	for _, spot := range l.Spots {
		if vehicle.Type <= spot.Size && spot.IsFree {
			return spot, nil
		}
	}
	return nil, errors.New("no slot is free")
}

func (l *Lot) UnparkedVehicle(tikcetId int, rate feecalculator.RatePerUnit) (float64, error) {
	for _, ticket := range l.Tickets {
		if ticket.Id == tikcetId && !ticket.IsPaid {
			fee := ticket.PriceCalculator(rate)
			ticket.IsPaid = true
			ticket.Spot.RemoveVehicle()
			return fee, nil
		}
	}
	return 0, errors.New("invalid ticket or already paid")
}
