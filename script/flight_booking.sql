--// flights
CREATE TABLE "flights" (
  "id" varchar PRIMARY KEY,	--ID
  "name_flights" varchar(200) NOT NULL,	--name of flight
  "departure_airport" varchar(20) NOT NULL,	--departure_airport
  "departure_arrival" varchar(20) NOT NULL,	--departure_arrival
  "depart_date" timestamptz NOT NULL,	--flight_date
  "status" varchar(10) NOT NULL,	--status	(1: active, 0: not_active)
  "available_slot" int NOT NULL,	-- number of slot available
  "created_at" timestamptz NOT NULL DEFAULT 'now()',
  "updated_at" timestamptz NOT NULL DEFAULT 'now()'
);

--// customer table
CREATE TABLE "customers" (
  "id" varchar PRIMARY KEY,	--ID
  "role" int,	--Role (0: not_register, 1: registed, 2: Admin)
  "customer_name" varchar(200) NOT NULL,	--customer_name
  "email" varchar(200) NOT NULL,	--Email
  "phone_number" varchar(20) NOT NULL,	--SĐT
  "date_of_bith" varchar(20) NOT NULL,	-- Ngày sinh
  "identity_card" varchar(20) NOT NULL,	--identity_card
  "address" varchar(200) NOT NULL,	--address
  "membership_card"  varchar(20),	--membership_card
  "password" varchar(200),	--Password
  "status" int,	--status (0: inactive, 1: Active)
  "created_at" timestamptz NOT NULL DEFAULT 'now()',
  "updated_at" timestamptz NOT NULL DEFAULT 'now()'
);


--// Lưu thông tin booking
CREATE TABLE "bookings" (
  "id" varchar PRIMARY KEY,
  "customer_id" varchar NOT NULL,	--customer_id
  "flight_id" varchar NOT NULL,	--flight_id
  "flight_number" varchar(20) NOT NULL,	--number flight
  "booked_slot" int,	-- Số ghế booking
  "status" varchar(10) NOT NULL,	-- status  booking
  "booked_date" timestamp NOT NULL DEFAULT 'now()',
  "created_at" timestamptz NOT NULL DEFAULT 'now()',
  "updated_at" timestamptz NOT NULL DEFAULT 'now()'
);

ALTER TABLE "bookings" ADD FOREIGN KEY ("customer_id") REFERENCES "customers" ("id");

ALTER TABLE "bookings" ADD FOREIGN KEY ("flight_id") REFERENCES "flights" ("id");
