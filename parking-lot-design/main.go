package main

type Vehicle struct {
	LicenseNumber string
	VehicleType   VehicleType
	ParkingTicket ParkingTicket
	PaymentInfo   PaymentInfo
}

type ParkingSpace struct {
	SpaceID          int
	IsFree           bool
	ParkingSpaceType ParkingSpaceType
	Vehicle          Vehicle
	CostPerHour      float32
}

type ParkingDisplayBoard struct {
	FreeSpotsAvailableMap map[ParkingSpaceType]int
}

func (pdb *ParkingDisplayBoard) UpdateFreeSpotsAvailable(parkingSpaceType ParkingSpaceType, spaces int) {
	pdb.FreeSpotsAvailableMap[parkingSpaceType] = spaces
}

type ParkingFloor struct {
	FloorID             int
	IsFull              bool
	ParkingSpaces       []ParkingSpace
	ParkingDisplayBoard ParkingDisplayBoard
}

type ParkingLot struct {
	ParkingLotName string
	Address        Address
	ParkingFloors  []ParkingFloor
	EntranceGates  []EntranceGate
	ExistGates     []ExitGate
}

func (c *ParkingLot) IsParkingSpaceAvailableForVehicle(vehicle Vehicle) bool {
	return true
}

func (c *ParkingLot) UpdateParkingAttendant(parkingAttendant ParkingAttendant, gateId int) bool {
	return true
}
