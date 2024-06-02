package main

import "time"

type PaymentInfo struct {
	Amount        float64
	PaymentDate   time.Time
	TransactionID int
	ParkingTicket ParkingTicket
	PaymentStatus PaymentStatus
}

type Payment struct{}

func (p *Payment) MakePayment(parkingTicket ParkingTicket, paymentType PaymentType) PaymentInfo
