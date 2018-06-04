package dao

import (
	"os"
	"context"
	"cloud.google.com/go/datastore"
	"app/models/dto"
	"log"
	"time"
)

func CallProcessDao()[]dto.Booking{
	ctx:=context.Background()
	projectID := os.Getenv("projectId")
	client, _ := datastore.NewClient(ctx, projectID)
	booking:=[]dto.Booking{}
	queryLogin:=datastore.NewQuery("booking").Namespace("TwilioAiko")
	if _,err:=client.GetAll(ctx,queryLogin,&booking);err!=nil{
		log.Printf("Don't find data")
	}

	return booking
}

func DeleteBookingDao(key *datastore.Key)bool{
	ctx:=context.Background()
	client:=ConnectDataStore()
	err:=client.Delete(ctx,key)
	if(err!=nil){
		return false
	}
	return true
}

func InsertBookingLastDao(booking dto.Booking)bool{
	client:=ConnectDataStore()
	task := datastore.Key{}
	task.Namespace = os.Getenv("namespace")
	task.Kind = "booking_last"
	bookingLast:=dto.BookingLast{}
	bookingLast.Status = "Called"
	bookingLast.CounselorSupport=""
	bookingLast.UserClient = booking.UserClient
	bookingLast.TimeCall = ""
	bookingLast.MoneyPaid=""
	bookingLast.Note=booking.Note
	bookingLast.UpdatedAt = time.Now()
	bookingLast.CreatedAt = time.Now()
	_,err1:=client.Put(context.Background(),&task,&bookingLast)
	if(err1!=nil){
		return false
	}
	return true
}