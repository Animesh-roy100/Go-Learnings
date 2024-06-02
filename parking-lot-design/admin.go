package main

type Address struct {
	Country string
	State   string
	City    string
	PinCode int
	// Add others
}

type Account struct {
	Name     string
	Email    string
	Password string
	EmpId    string
	Address  Address
}

type Admin struct {
	Account Account
}

func (a *Admin) AddParkingFloor(parkingLot ParkingLot, parkingFloor ParkingFloor) bool {
	return true
}

func (a *Admin) AddParkingSpace(parkingFloor ParkingFloor, parkingSpace ParkingSpace) bool {
	return true
}

func (a *Admin) AddParkingDisplayBoard(parkingFloor ParkingFloor, parkingDisplayBoard ParkingDisplayBoard) bool {
	return true
}

type ParkingAttendant struct {
	Account        Account
	PaymentService Payment
}

func (pa *ParkingAttendant) ProcessVehicleEntry(vehicle Vehicle) {}

func (pa *ParkingAttendant) ProcessPayment(parkingType ParkingTicket, paymentType PaymentType) PaymentInfo {
	return PaymentInfo{}
}
