package fee

import "time"

type ParkingFeeCalculator interface {
	CalculateFee(entryTime, exitTime time.Time) float64
}
