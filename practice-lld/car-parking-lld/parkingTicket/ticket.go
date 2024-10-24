package parkingticket

import (
	feecalculator "car-parking/feeCalculator"
	parkingslot "car-parking/parkingSlot"
	"car-parking/vehicle"
	"time"
)

type Ticket struct {
	Id        int
	Spot      *parkingslot.Spot
	EntryTime time.Time
	ExitTime  time.Time
	Vehicle   vehicle.Vehicle
	IsPaid    bool
	FeeCharge float64
}

func (t *Ticket) PriceCalculator(rate feecalculator.RatePerUnit) float64 {
	t.FeeCharge = rate.CalculateFee(t.EntryTime, time.Now())
	return t.FeeCharge
}
