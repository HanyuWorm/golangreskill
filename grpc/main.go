package main

import (
	"flag"
	"fmt"
	booking_repo "mock-golang/grpc/booking-grpc/repository"
	booking_handler "mock-golang/grpc/booking-grpc/service"
	customer_repo "mock-golang/grpc/customer-grpc/repository"
	customer_handler "mock-golang/grpc/customer-grpc/service"
	flight_repo "mock-golang/grpc/flight-grpc/repository"
	flight_handler "mock-golang/grpc/flight-grpc/service"
	"mock-golang/helper"
	"mock-golang/intercepter"
	"mock-golang/protobuf"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	configFile = flag.String("config-file", "../helper/config.yml", "Location of config file")
	port       = flag.Int("port", 9112, "Port grpc")
)

func init() {
	flag.Parse()
}

func main() {
	err := helper.AutoBindConfig(*configFile)
	if err != nil {
		panic(err)
	}

	listen, err := net.Listen("tcp", fmt.Sprintf(":%v", *port))
	if err != nil {
		panic(err)
	}

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			intercepter.UnaryServerLoggingIntercepter(logger),
		)),
	)

	reflection.Register(s)

	// Initial customer repository START
	customerRepository, err := customer_repo.NewDBManager()
	if err != nil {
		panic(err)
	}

	h, err := customer_handler.NewCustomerHandler(customerRepository)
	if err != nil {
		panic(err)
	}
	protobuf.RegisterRPCCustomerServer(s, h)
	// Initial customer repository END

	// Initial Flight repository START
	flightRepository, errFlight := flight_repo.NewDBManager()
	if errFlight != nil {
		panic(errFlight)
	}

	hFlight, errFlight := flight_handler.NewFlightHandler(flightRepository)
	if errFlight != nil {
		panic(errFlight)
	}
	protobuf.RegisterRPCFlightServer(s, hFlight)
	// Initial Flight repository END

	// Initial Booking repository START
	bookingRepository, errBooking := booking_repo.NewDBManager()
	if errBooking != nil {
		panic(errBooking)
	}

	hBooking, errBooking := booking_handler.NewBookingHandler(bookingRepository)
	if errBooking != nil {
		panic(errBooking)
	}
	protobuf.RegisterRPCBookingServer(s, hBooking)
	// Initial Booking repository END

	fmt.Printf("Listen at port: %v\n", *port)

	s.Serve(listen)
}
