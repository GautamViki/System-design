package main

import (
	"car-parking/fee"
	parkinglot "car-parking/parking-lot"
	"car-parking/vehicle"
	"fmt"
	"time"
)

func main() {
	parkingLot := parkinglot.NewParkingLot(10, 10, 5)
	// Create a vehicle
	vehicle := &vehicle.Vehicle{
		LicencePlate: "ABC123",
		VehicleType:  vehicle.Medium,
	}

	// Park the vehicle
	ticket, err := parkingLot.ParkingVehicle(vehicle)
	if err != nil {
		fmt.Println("Error parking vehicle:", err)
		return
	}
	fmt.Printf("Vehicle parked. Ticket ID: %d\n", ticket.TicketId)

	// Simulate vehicle leaving after some time
	time.Sleep(5 * time.Second)

	// Calculate parking fee
	feeCalculator := &fee.HourlyFeeCalculator{HourlyRate: 10}
	fee, err := parkingLot.RetrieveVehicle(ticket.TicketId, feeCalculator)
	if err != nil {
		fmt.Println("Error retrieving vehicle:", err)
		return
	}
	fmt.Printf("Parking fee: $%.2f\n", fee)
}
