package main

type PaymentType int
type ParkingSpaceType int
type VehicleType int
type ParkingTicketStatus int
type PaymentStatus int

const (
	Car VehicleType = iota
	Bike
	Truck
)

const (
	Cash PaymentType = iota
	CreditCard
	DebitCard
	UPI
)

const (
	BikeParking ParkingSpaceType = iota
	CarParking
	TruckParking
)

const (
	Paid ParkingTicketStatus = iota
	Active
)

const (
	Unpaid PaymentStatus = iota
	Pending
	Completed
	Declined
	Cancelled
	Refunded
)
