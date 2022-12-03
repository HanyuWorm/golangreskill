package flight_repo

import (
	"context"
	"mock-golang/database"
	flight_model "mock-golang/grpc/flight-grpc/model"
	flight_request "mock-golang/grpc/flight-grpc/request"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

//Embeded struct

type FlightRepository interface {
	FindById(ctx context.Context, id uuid.UUID) (*flight_model.Flight, error)
	CreateFlight(ctx context.Context, model *flight_model.Flight) (*flight_model.Flight, error)
	UpdateFlight(ctx context.Context, model *flight_model.Flight) (*flight_model.Flight, error)
	SearchFlight(ctx context.Context, req *flight_request.SearchFlightRequest) ([]*flight_model.Flight, error)
}

type dbmanager struct {
	*gorm.DB
}

func NewDBManager() (FlightRepository, error) {
	db, err := database.NewGormDB()
	if err != nil {
		return nil, err
	}

	db = db.Debug()

	err = db.AutoMigrate(
		&flight_model.Flight{},
	)

	if err != nil {
		return nil, err
	}

	return &dbmanager{db}, nil
}

func (m *dbmanager) FindById(ctx context.Context, id uuid.UUID) (*flight_model.Flight, error) {
	res := flight_model.Flight{}
	if err := m.Where(&flight_model.Flight{Id: id}).First(&res).Error; err != nil {
		return nil, err
	}

	return &res, nil
}

func (m *dbmanager) CreateFlight(ctx context.Context, model *flight_model.Flight) (*flight_model.Flight, error) {
	if err := m.Create(model).Error; err != nil {
		return nil, err
	}

	return model, nil
}

func (m *dbmanager) UpdateFlight(ctx context.Context, model *flight_model.Flight) (*flight_model.Flight, error) {
	if err := m.Where(&flight_model.Flight{Id: model.Id}).Updates(model).Error; err != nil {
		return nil, err
	}

	return model, nil
}

func (m *dbmanager) SearchFlight(ctx context.Context, req *flight_request.SearchFlightRequest) ([]*flight_model.Flight, error) {
	flights := []*flight_model.Flight{}

	sbWhere := " 1=1 "
	params := []interface{}{}
	if len(strings.TrimSpace(req.Id)) > 0 {
		sbWhere += " AND Id = ? "
		params = append(params, req.Id)
	}
	if len(strings.TrimSpace(req.Name)) > 0 {
		sbWhere += " AND Name = ? "
		params = append(params, req.Name)
	}
	if len(strings.TrimSpace(req.From)) > 0 {
		sbWhere += " AND From = ? "
		params = append(params, req.From)
	}
	if len(strings.TrimSpace(req.To)) > 0 {
		sbWhere += " AND To = ? "
		params = append(params, req.To)
	}
	if !req.FromDate.IsZero() {
		sbWhere += " AND depart_date >= ? "
		params = append(params, req.FromDate)
	}
	if !req.ToDate.IsZero() {
		sbWhere += " AND depart_date <= ? "
		params = append(params, req.ToDate)
	}

	if err := m.Where(sbWhere, params...).Find(&flights).Error; err != nil {
		return nil, err
	}

	return flights, nil
}
