# Golang reskill project

This is an example and `simple` project to mock flight-management system.

It runs under microservices architecture with 3 services:

- Booking: Manage reserved ticket for users and flights.
- User: Manage users.
- Flight: Manage flights

## Project structure

### Helper

- Configuration: manage configuration
- Database: contains database connection initialization

### User

- Located in folder `/customer`
- Restful API served:

POST `/customer` - register new customer

PUT `/customer/:username` - update customer data by id

POST `/customer/viewBookingHistory` - update customer data

POST `/customer/searchBooking` - Search Booking data

POST `/customer/changePassword` - Change Password

- gRPC served:

Same with rest api

### Flight

- Located in folder `/flight`
- Restful API served:

POST `/flight` - Create Flight

GET `/flight/:id` - Get flight by id

PUT `/flight/:id` - Update Flight



- gRPC served:

Same with rest api

### Booking

- Located in folder `booking`
- Restful API served:

POST `/booking` - Create Booking

GET `/booking/guest` - Get list of user's reserved bookings

GET `/booking/cancel` - Cancel booking

- gRPC served:

Same with rest api
## Usage

- This project's purpose is for user viewpoint. Therefore, their actions should be:
- View flights
- View personal data
- Reserve flights


###  GUIDE Run API
1. Run script create table
   mock-golang\script
   FileName: flight_booking.sql

2. Update information DB
   mock-golang\helper
   FileName: config.yml

3. run server grpc (port 2222)
   Ex: mock-golang\grpc
   cmd: go run main.go

4. Run server api (port 3333)
   mock-golang\api
   cmd: go run main.go