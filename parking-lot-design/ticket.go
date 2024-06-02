package main

import "time"

type ParkingTicket struct {
	TicketID             int
	FloorID              int
	SpaceID              int
	VehicleEntryDateTime time.Time
	VehicleExitDateTime  time.Time
	ParkingSpaceType     ParkingSpaceType
	TotalCost            float64
	ParkingTicketStatus  ParkingTicketStatus
}

func (pt *ParkingTicket) UpdateTotalCost() {}

func (pt *ParkingTicket) UpdateVehicleExitTime(vehicleExitDateTime time.Time) {
	pt.VehicleExitDateTime = vehicleExitDateTime
	pt.UpdateTotalCost()
}
