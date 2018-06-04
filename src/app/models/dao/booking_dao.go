package dao

import (
	"app/models/form"
	"app/models/dto"
	"github.com/google/uuid"
	"time"
	"os"
	"context"
	"cloud.google.com/go/datastore"
	"log"
	"google.golang.org/api/iterator"
)

func BookingDao(booking form.BookingForm,userId int)error{
	bookingDto:=dto.Booking{}
	bookingDto.BookingId = uuid.New().String()
	bookingDto.UserClient = userId;
	bookingDto.Note = booking.Note
	bookingDto.BookingDate = booking.Date
	bookingDto.UpdatedAt = time.Now()
	bookingDto.CreatedAt=time.Now()
	// update data to datastore
	ctx:=context.Background()
	taskKey := datastore.Key{}
	taskKey.Namespace = os.Getenv("namespace")
	taskKey.Kind = "booking"
	client:=ConnectDataStore()
	_,err:= client.Put(ctx, &taskKey, &bookingDto)
	return err
}

func EditBookingDao(booking form.BookingForm)error{
	bookingDto:=dto.Booking{}
	client:=ConnectDataStore()
	tx,err:=client.NewTransaction(context.Background())
	if err!=nil{
		log.Fatalf("client.NewTransaction: %v",err)
	}

	bookingQuery:=datastore.NewQuery("booking").Namespace("TwilioAiko").Filter("booking_id=",booking.BookingId)
	it:=client.Run(context.Background(),bookingQuery)
	for {
		_, err := it.Next(&bookingDto)
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Error fetching next task: %v", err)
		}
	}
	bookingDto.Note = booking.Note
	bookingDto.BookingDate = booking.Date
	bookingDto.UpdatedAt = time.Now()
	if _,err:=tx.Put(bookingDto.K,&bookingDto);err!=nil{
		log.Fatalf("tx.Put:%v",err)
	}
	if _,err:=tx.Commit();err!=nil{
		log.Printf("tx.Commit:%v",err)
	}
	return err
}