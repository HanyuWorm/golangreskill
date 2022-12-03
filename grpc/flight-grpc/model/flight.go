package flight_model

import (
	"time"

	"mock-golang/protobuf"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Flight struct {
	Id               uuid.UUID `gorm:"type:uuid;primaryKey"`
	NameFlight       string    `gorm:"column:flights"`
	DepartureAirport string    `gorm:"column:departure_airport"`
	DepartureArrival string    `gorm:"column:departure_arrival"`
	DepartDate       time.Time `gorm:"column:depart_date"`
	Status           string    `gorm:"column:status"`
	AvailableSlot    int32     `gorm:"column:available_slot"`
	CreatedAt        time.Time `gorm:"column:created_at"`
	UpdatedAt        time.Time `gorm:"column:updated_at"`
}

func (in *Flight) ToResponse() *protobuf.Flight {
	res := &protobuf.Flight{
		Id:            in.Id.String(),
		Name:          in.NameFlight,
		From:          in.DepartureAirport,
		To:            in.DepartureArrival,
		DepartDate:    timestamppb.New(in.DepartDate),
		Status:        in.Status,
		AvailableSlot: in.AvailableSlot,
		CreatedAt:     timestamppb.New(in.CreatedAt),
		UpdatedAt:     timestamppb.New(in.UpdatedAt),
	}

	return res
}
