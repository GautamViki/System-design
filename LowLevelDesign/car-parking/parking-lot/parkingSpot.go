package parkinglot

import (
	"car-parking/vehicle"
	"errors"
)

type ParkingSpot struct {
	SpotId    int
	SpotSize  vehicle.VehicleType
	IsFree    bool
	ParkedCar vehicle.Vehicle
}

func (ps *ParkingSpot) ParkVehicle(v *vehicle.Vehicle) error {
	if ps.IsFree && ps.SpotSize >= v.VehicleType {
		ps.ParkedCar = *v
		ps.IsFree = false
		return nil
	}
	return errors.New("spot not available or vehicle too large")
}

func (ps *ParkingSpot) RemoveVehicle() {
	ps.ParkedCar = vehicle.Vehicle{}
	ps.IsFree = true
}
