package booking_request

import (
	"time"

	"github.com/google/uuid"
)

type BookingRequest struct {
	Id uuid.UUID
}

type SearchBookingRequest struct {
	Id         string
	CustomerId string
	FlightId   string
	Code       string
	Status     string
	FromDate   time.Time
	ToDate     time.Time
}
