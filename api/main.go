package main

import (
	booking_handler "mock-golang/api/booking-api/service"
	customer_handler "mock-golang/api/customer-api/service"
	flight_handler "mock-golang/api/flight-api/service"
	"mock-golang/middleware"
	"mock-golang/protobuf"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	//Create grpc client connect
	conn, err := grpc.Dial(":9112", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	//Singleton
	customerClient := protobuf.NewRPCCustomerClient(conn)
	bookingClient := protobuf.NewRPCBookingClient(conn)
	flightClient := protobuf.NewRPCFlightClient(conn)

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	//Handler for GIN Gonic
	hCustomer := customer_handler.NewCustomerHandler(customerClient)
	hFlight := flight_handler.NewFlightHandler(flightClient)
	hBooking := booking_handler.NewBookingHandler(bookingClient, customerClient, flightClient)
	os.Setenv("GIN_MODE", "debug")
	g := gin.Default()
	g.Use(middleware.LoggingMiddleware(logger))

	//Create routes
	gr := g.Group("/v1/api")

	// API Customer
	gr.POST("/customer", hCustomer.CreateCustomer)
	gr.PUT("/customer", hCustomer.UpdateCustomer)
	gr.POST("/customer/changePassword", hCustomer.ChangePassword)
	gr.POST("/customer/viewBookingHistory", hBooking.BookingHistory)
	gr.POST("/customer/searchBooking", hBooking.SearchBooking)

	// API Booking
	gr.POST("/booking", hBooking.CustomerBooking)
	gr.POST("/booking/guest", hBooking.GuestBooking)
	gr.POST("/booking/cancel", hBooking.CancelBooking)

	// API Flight
	gr.POST("/flight", hFlight.CreateFlight)
	gr.PUT("/flight", hFlight.UpdateFlight)
	gr.GET("/flight/search", hFlight.SearchFlight)
	gr.GET("/flight/:id", hFlight.SearchFlightById)

	//Listen and serve
	http.ListenAndServe(":8080", g)
}
