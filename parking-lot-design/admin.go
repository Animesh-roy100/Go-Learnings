package main

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
