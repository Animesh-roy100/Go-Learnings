package main

type Gate struct {
	GateId           int
	ParkingAttendant ParkingAttendant
}

// Inheritance
type EntranceGate struct {
	Gate Gate
}

func (eg *EntranceGate) GetParkingTicket(vehicle Vehicle) ParkingTicket {
	return ParkingTicket{}
}

type ExitGate struct {
	Gate Gate
}

func (eg *ExitGate) PayForParking(parkingTicket ParkingTicket, paymentType PaymentType) ParkingTicket {
	return ParkingTicket{}
}
