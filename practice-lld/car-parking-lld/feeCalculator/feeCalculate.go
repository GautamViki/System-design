package feecalculator

import "time"

type feeCalculator interface {
	CalculateFee(entryTime, exitTime time.Time) float64
}
