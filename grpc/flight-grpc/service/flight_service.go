package flight_handler

import (
	"context"
	"database/sql"
	flight_model "mock-golang/grpc/flight-grpc/model"
	flight_repo "mock-golang/grpc/flight-grpc/repository"
	flight_request "mock-golang/grpc/flight-grpc/request"
	"mock-golang/protobuf"
	"sync"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type FlightHandler struct {
	protobuf.UnimplementedRPCFlightServer
	flightRepository flight_repo.FlightRepository
	mu               *sync.Mutex
}

func NewFlightHandler(flightRepository flight_repo.FlightRepository) (*FlightHandler, error) {
	return &FlightHandler{
		flightRepository: flightRepository,
		mu:               &sync.Mutex{},
	}, nil
}

func (h *FlightHandler) FindById(ctx context.Context, in *protobuf.FlightParamId) (*protobuf.Flight, error) {
	flight, err := h.flightRepository.FindById(ctx, uuid.MustParse(in.Id))

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return flight.ToResponse(), nil
}

func (h *FlightHandler) CreateFlight(ctx context.Context, in *protobuf.Flight) (*protobuf.Flight, error) {
	req := &flight_model.Flight{
		Id:               uuid.New(),
		NameFlight:       in.Name,
		DepartureAirport: in.From,
		DepartureArrival: in.To,
		DepartDate:       in.DepartDate.AsTime(),
		Status:           in.Status,
		AvailableSlot:    in.AvailableSlot,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}
	flight, err := h.flightRepository.CreateFlight(ctx, req)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return flight.ToResponse(), nil
}

func (h *FlightHandler) UpdateFlight(ctx context.Context, in *protobuf.Flight) (*protobuf.Flight, error) {
	flightIn, err := h.flightRepository.FindById(ctx, uuid.MustParse(in.Id))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, err
	}

	if in.Name != "" {
		flightIn.NameFlight = in.Name
	}

	if in.From != "" {
		flightIn.DepartureAirport = in.From
	}

	if in.To != "" {
		flightIn.DepartureArrival = in.To
	}

	if in.DepartDate != nil {
		flightIn.DepartDate = in.DepartDate.AsTime()
	}

	if in.AvailableSlot > 0 {
		flightIn.AvailableSlot = in.AvailableSlot
	}

	if in.Status != "" {
		flightIn.Status = in.Status
	}

	flightIn.UpdatedAt = time.Now()

	flight, err := h.flightRepository.UpdateFlight(ctx, flightIn)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return flight.ToResponse(), nil
}

func (h *FlightHandler) SearchFlight(ctx context.Context, in *protobuf.SearchFlightRequest) (*protobuf.SearchFlightResponse, error) {
	flights, err := h.flightRepository.SearchFlight(ctx, &flight_request.SearchFlightRequest{
		Id:   in.Id,
		Name: in.Name,
		From: in.From,
		To:   in.To,
		// FromDate: in.FromDate.AsTime(),
		// ToDate:   in.ToDate.AsTime(),
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, err
	}

	pRes := &protobuf.SearchFlightResponse{
		Flight: []*protobuf.Flight{},
	}

	for _, flight := range flights {
		pRes.Flight = append(pRes.Flight, flight.ToResponse())
	}

	if err != nil {
		return nil, err
	}

	return pRes, nil
}
