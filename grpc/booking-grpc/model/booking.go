package booking_model

import (
	customer_model "mock-golang/grpc/customer-grpc/model"
	flight_model "mock-golang/grpc/flight-grpc/model"
	"time"

	"mock-golang/protobuf"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Booking struct {
	Id         uuid.UUID                `gorm:"type:uuid;primaryKey"`
	CustomerId string                   `gorm:"column:customer_id"`
	FlightId   string                   `gorm:"column:flight_id"`
	Code       string                   `gorm:"column:flight_number"`
	BookedSlot int32                    `gorm:"column:booked_slot"`
	BookedDate time.Time                `gorm:"column:booked_date"`
	Status     string                   `gorm:"column:status"`
	CreatedAt  time.Time                `gorm:"column:created_at"`
	UpdatedAt  time.Time                `gorm:"column:updated_at"`
	Customer   *customer_model.Customer `gorm:"foreignKey:customer_id;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Flight     *flight_model.Flight     `gorm:"foreignKey:flight_id;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (in *Booking) ToResponse() *protobuf.Booking {
	res := &protobuf.Booking{
		Id:         in.Id.String(),
		CustomerId: in.CustomerId,
		FlightId:   in.FlightId,
		Code:       in.Code,
		BookedSlot: in.BookedSlot,
		BookedDate: timestamppb.New(in.BookedDate),
		Status:     in.Status,
		CreatedAt:  timestamppb.New(in.CreatedAt),
		UpdatedAt:  timestamppb.New(in.UpdatedAt),
		Customer: &protobuf.CustomerDTO{
			Id:             in.Customer.Id.String(),
			Role:           in.Customer.Role,
			Name:           in.Customer.Name,
			Email:          in.Customer.Email,
			PhoneNumber:    in.Customer.PhoneNumber,
			DateOfBith:     in.Customer.DateOfBith,
			IdentityCard:   in.Customer.IdentityCard,
			Address:        in.Customer.Address,
			MembershipCard: in.Customer.MembershipCard,
			Password:       in.Customer.Password,
			Status:         in.Customer.Status,
			CreatedAt:      timestamppb.New(in.Customer.CreatedAt),
			UpdatedAt:      timestamppb.New(in.Customer.UpdatedAt),
		},
		Flight: &protobuf.FlightDTO{
			Id:            in.Flight.Id.String(),
			Name:          in.Flight.NameFlight,
			From:          in.Flight.DepartureAirport,
			To:            in.Flight.DepartureArrival,
			DepartDate:    timestamppb.New(in.Flight.DepartDate),
			Status:        in.Flight.Status,
			AvailableSlot: in.Flight.AvailableSlot,
			CreatedAt:     timestamppb.New(in.Flight.CreatedAt),
			UpdatedAt:     timestamppb.New(in.Flight.UpdatedAt),
		},
	}

	return res
}

func (in *Booking) ToResponseForCreate() *protobuf.Booking {
	res := &protobuf.Booking{
		Id:         in.Id.String(),
		CustomerId: in.CustomerId,
		FlightId:   in.FlightId,
		Code:       in.Code,
		BookedSlot: in.BookedSlot,
		BookedDate: timestamppb.New(in.BookedDate),
		Status:     in.Status,
		CreatedAt:  timestamppb.New(in.CreatedAt),
		UpdatedAt:  timestamppb.New(in.UpdatedAt),
	}

	return res
}
