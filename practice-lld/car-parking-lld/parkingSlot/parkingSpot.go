package parkingslot

import (
	"car-parking/vehicle"
	"errors"
)

type Spot struct {
	SpotId        int
	Size          vehicle.VehicleType
	IsFree        bool
	ParkedVehicle vehicle.Vehicle
}

func (s *Spot) ParkCar(vehicle *vehicle.Vehicle) error {
	if s.IsFree && s.Size >= vehicle.Type {
		s.Size = vehicle.Type
		s.IsFree = true
		s.ParkedVehicle = *vehicle
		return nil
	}
	return errors.New("spot not available or vehicle too large")
}

func (s *Spot) RemoveVehicle() {
	s.ParkedVehicle = vehicle.Vehicle{}
	s.IsFree = true
}
