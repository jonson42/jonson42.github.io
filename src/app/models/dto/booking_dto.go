package dto

import (
	"time"
	"cloud.google.com/go/datastore"
)

type Booking struct {
	K *datastore.Key `datastore:"__key__"`
	BookingId string `datastore:"booking_id"`
	BookingDate string `datastore:"booking_date"`
	ReceiveBy string `datastore:"receive_by"`
	UserClient int `datastore:"user_client"`
	Note string `datastore:"note"`
	CreatedAt time.Time `datastore:"created_at"`
	UpdatedAt time.Time `datastore:"updated_at"`
}