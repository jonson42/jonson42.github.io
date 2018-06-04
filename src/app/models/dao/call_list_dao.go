package dao
import(
	"app/models/dto"
	"os"
	"context"
	"cloud.google.com/go/datastore"
	"log"
)
func GetListDao(userId int)[]dto.BookingLast{
	booking:=[]dto.Booking{}
	bookingLast:=[]dto.BookingLast{}
	bookingResult:=[]dto.BookingLast{}
	projectID := os.Getenv("projectId")
	client, _ := datastore.NewClient(context.Background(), projectID)
	queryBooking:=datastore.NewQuery("booking").Namespace("TwilioAiko").Filter("user_client=",userId)
	if _,err:=client.GetAll(context.Background(),queryBooking,&booking);err!=nil{
		log.Printf("Don't find data")
	}
	queryBookingLast:=datastore.NewQuery("booking_last").Namespace("TwilioAiko").Filter("user_client",userId)
	if _,err:=client.GetAll(context.Background(),queryBookingLast,&bookingLast);err!=nil{
		log.Printf("Don't find data")
	}
	bookingResult=ChangeToListBooking(booking,bookingLast)
	return bookingResult
}

func ChangeToListBooking(booking []dto.Booking,bookingLast []dto.BookingLast)[]dto.BookingLast{
	for _,i:=range booking{
		bookTemp:=dto.BookingLast{}
		bookTemp.Status="Open"
		bookTemp.UserClient = i.UserClient
		bookTemp.DateTime = i.BookingDate
		bookTemp.Note = i.Note
		bookTemp.UpdatedAt = i.UpdatedAt
		bookTemp.CreatedAt = i.CreatedAt
		bookingLast=append(bookingLast, bookTemp)
	}
	return bookingLast
}
