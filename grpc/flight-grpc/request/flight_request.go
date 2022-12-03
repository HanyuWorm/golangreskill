package flight_request

import (
	"time"

	"github.com/google/uuid"
)

type FlightRequest struct {
	Id uuid.UUID
}

type SearchFlightRequest struct {
	Id       string
	Name     string
	From     string
	To       string
	FromDate time.Time
	ToDate   time.Time
}
