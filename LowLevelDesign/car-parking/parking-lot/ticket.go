package parkinglot

import (
	"car-parking/fee"
	"car-parking/vehicle"
	"time"
)

type ParkingTicket struct {
	TicketId   int
	Vehicle    *vehicle.Vehicle
	Spot       *ParkingSpot
	EntryTime  time.Time
	ExitTime   time.Time
	FeeCharged float64
	IsPaid     bool
}

func (pt *ParkingTicket) CalculateFee(feeCalculator fee.ParkingFeeCalculator) float64 {
	pt.ExitTime = time.Now()
	pt.FeeCharged = feeCalculator.CalculateFee(pt.EntryTime, pt.ExitTime)
	return pt.FeeCharged
}
