package form

import "cloud.google.com/go/datastore"

type BookingForm struct {
	K *datastore.Key `datastore:"__key__"`
	Date string `form:"date"`
	BookingId string `form:"bookingId"`
	//User int `form:"userId`
	Note string `form:"note"`
}
