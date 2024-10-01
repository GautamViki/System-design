package fee

import (
	"time"
)

type HourlyFeeCalculator struct {
	HourlyRate float64
}

func (h *HourlyFeeCalculator) CalculateFee(entryTime, exitTime time.Time) float64 {
	duration := exitTime.Sub(entryTime)
	totalHours := duration.Seconds()
	return totalHours * h.HourlyRate
}
