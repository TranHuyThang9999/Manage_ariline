CREATE TABLE users (
  user_id text PRIMARY KEY,
  user_name TEXT NOT NULL,
  phone_number TEXT NOT NULL,
  email TEXT NOT NULL UNIQUE,
  password TEXT NOT NULL,
  address TEXT NOT NULL,
  date_birth text NOT NULL,
  number_cmnd TEXT NOT NULL,
  nationality TEXT NOT NULL,
  language TEXT NOT NULL,
  amount FLOAT NOT NULL CHECK(amount >= 0)
);



CREATE TABLE flights (
  flight_id text PRIMARY KEY,
  name_airline TEXT NOT NULL,
  destination TEXT NOT NULL,
  departure TEXT NOT NULL,
  destination_time text NOT NULL,
  departure_time text NOT NULL,
  remaining_seats INT,
  ticket_type TEXT NOT NULL,
  fare FLOAT NOT NULL CHECK(fare >= 0),
  status TEXT NOT NULL,
  name_flight text
);


CREATE TABLE bookings (
  booking_id text PRIMARY KEY,
  user_id text NOT NULL REFERENCES users(user_id), -- token header
  flight_id text NOT NULL REFERENCES flights(flight_id),
  number_of_seats INT,
  amount FLOAT NOT NULL CHECK(amount >= 0),
  FOREIGN KEY (user_id) REFERENCES users(user_id),
  FOREIGN KEY (flight_id) REFERENCES flights(flight_id),
  
  
    user_name TEXT NOT NULL,
    phone_number TEXT NOT NULL,
    address TEXT NOT NULL,
    number_cmnd TEXT NOT NULL,
    nationality TEXT NOT NULL,
  language TEXT NOT NULL,
    name_airline TEXT NOT NULL,
  destination TEXT NOT NULL,
  departure TEXT NOT NULL,
  destination_time text NOT NULL,
  departure_time text NOT NULL,
  remaining_seats INT,
  ticket_type TEXT NOT NULL,
  fare FLOAT NOT NULL CHECK(fare >= 0),
  status TEXT NOT NULL,
  name_flight text,
register_time text
);

CREATE TABLE payments (
  payment_id text PRIMARY KEY,
  booking_id text NOT NULL REFERENCES bookings(booking_id),
  payment_time TIMESTAMP NOT NULL DEFAULT NOW(),
  amount FLOAT NOT NULL CHECK(amount >= 0),
  FOREIGN KEY (booking_id) REFERENCES bookings(booking_id)
);

create table admins(
  user_id text PRIMARY KEY,
  user_name TEXT NOT NULL,
  phone_number TEXT NOT NULL,
  email TEXT ,
  password TEXT NOT NULL,
  address TEXT NOT NULL,
  date_birth text NOT NULL,
  number_cmnd TEXT NOT NULL,
  nationality TEXT NOT NULL,
  language TEXT NOT NULL
);



 SELECT * FROM "users" WHERE phone_number = 'fer' ORDER BY "users"."user_id" LIMIT 1
SELECT * FROM "admins" WHERE phone_number = '01234' ORDER BY "admins"."user_id" LIMIT 1

select *from users

select *from admins

SELECT * FROM "users" WHERE phone_number = 'string2' ORDER BY "users"."user_id" LIMIT 1

select *from flights

--select *from bookings

DELETE FROM bookings;

 SELECT * FROM "bookings" WHERE phone_number = '' AND booking_id = '85af4fed-a9f0-4503-b020-9d6f7fb03aa5


