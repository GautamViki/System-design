package feecalculator

import (
	"time"
)

type RatePerUnit struct {
	RateCost float64
}

func (r *RatePerUnit) CalculateFee(entryTime, exitTime time.Time) float64 {
	duration := exitTime.Sub(entryTime)
	totalUnitTime := duration.Seconds()
	return totalUnitTime * r.RateCost
}
