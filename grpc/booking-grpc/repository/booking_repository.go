package booking_repo

import (
	"context"
	"mock-golang/database"
	booking_model "mock-golang/grpc/booking-grpc/model"
	booking_request "mock-golang/grpc/booking-grpc/request"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

//Embeded struct

type BookingRepository interface {
	FindById(ctx context.Context, id uuid.UUID) (*booking_model.Booking, error)
	CreateBooking(ctx context.Context, model *booking_model.Booking) (*booking_model.Booking, error)
	UpdateBooking(ctx context.Context, model *booking_model.Booking) (*booking_model.Booking, error)
	SearchBooking(ctx context.Context, model *booking_request.SearchBookingRequest) ([]*booking_model.Booking, error)
}

type dbmanager struct {
	*gorm.DB
}

func NewDBManager() (BookingRepository, error) {
	db, err := database.NewGormDB()
	if err != nil {
		return nil, err
	}

	db = db.Debug()

	err = db.AutoMigrate(
		&booking_model.Booking{},
	)

	if err != nil {
		return nil, err
	}

	return &dbmanager{db}, nil
}

func (m *dbmanager) FindById(ctx context.Context, id uuid.UUID) (*booking_model.Booking, error) {
	res := booking_model.Booking{}
	if err := m.Where(&booking_model.Booking{Id: id}).Preload("Customer").Preload("Flight").First(&res).Error; err != nil {
		return nil, err
	}

	return &res, nil
}

func (m *dbmanager) CreateBooking(ctx context.Context, model *booking_model.Booking) (*booking_model.Booking, error) {
	if err := m.Create(model).Error; err != nil {
		return nil, err
	}

	return model, nil
}

func (m *dbmanager) UpdateBooking(ctx context.Context, model *booking_model.Booking) (*booking_model.Booking, error) {
	if err := m.Where(&booking_model.Booking{Id: model.Id}).Updates(model).Error; err != nil {
		return nil, err
	}

	return model, nil
}

func (m *dbmanager) SearchBooking(ctx context.Context, req *booking_request.SearchBookingRequest) ([]*booking_model.Booking, error) {
	bookings := []*booking_model.Booking{}

	sbWhere := " 1=1 "
	params := []interface{}{}
	if len(strings.TrimSpace(req.Id)) > 0 {
		sbWhere += " AND Id = ? "
		params = append(params, req.Id)
	}
	if len(strings.TrimSpace(req.CustomerId)) > 0 {
		sbWhere += " AND customer_id = ? "
		params = append(params, req.CustomerId)
	}
	if len(strings.TrimSpace(req.FlightId)) > 0 {
		sbWhere += " AND flight_id = ? "
		params = append(params, req.FlightId)
	}
	if len(strings.TrimSpace(req.Code)) > 0 {
		sbWhere += " AND Code = ? "
		params = append(params, req.Code)
	}
	if !req.FromDate.IsZero() {
		sbWhere += " AND booked_date >= ? "
		params = append(params, req.FromDate)
	}
	if !req.ToDate.IsZero() {
		sbWhere += " AND booked_date <= ? "
		params = append(params, req.ToDate)
	}
	if len(strings.TrimSpace(req.Status)) > 0 {
		sbWhere += " AND Status = ? "
		params = append(params, req.Status)
	}

	if err := m.Where(sbWhere, params...).Preload("Customer").Preload("Flight").Find(&bookings).Error; err != nil {
		return nil, err
	}

	return bookings, nil
}
