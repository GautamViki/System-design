package main

import (
	feecalculator "car-parking/feeCalculator"
	parkinglot "car-parking/parkingLot"
	"car-parking/vehicle"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	lot := parkinglot.NewParkingLot(10, 10, 5)

	vehicle := vehicle.Vehicle{
		VehicleId: "123A",
		Type:      vehicle.Large,
	}

	ticket, err := lot.ParkedVehicle(&vehicle)
	if err != nil {
		fmt.Println("Error parking vehicle:", err.Error())
		return
	}
	fmt.Printf("Vehicle parked. Ticket ID: %d\n", ticket.Id)
	time.Sleep(5 * time.Second)
	fee := feecalculator.RatePerUnit{
		RateCost: 10,
	}
	cost, err := lot.UnparkedVehicle(ticket.Id, fee)
	if err != nil {
		fmt.Println("Error retrieving vehicle:", err.Error())
		return
	}

	fmt.Printf("Parking fee: $%.2f\n", cost)
}
