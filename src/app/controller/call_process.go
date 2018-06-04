package controller

import (
	"github.com/gin-gonic/gin"
	"time"
	"math/rand"
	"app/models/dao"
	"app/models/dto"
	"fmt"
	"strings"
	"net/http"
)

func CallProcess(c *gin.Context){
	go BackGroundProcessCall()
}
func BackGroundProcessCall(){
	c := time.Tick(1 * time.Minute)
	for now := range c {
		rand.Seed(time.Now().Unix())
		booking:=[]dto.Booking{}
		println(now.String())
		booking=dao.CallProcessDao()
		for i:=0;i<len(booking);i++{
			t:=strings.Replace(time.Now().String()[0:16]," ","T",-1)
			if(booking[i].BookingDate==t){
				//check number of
				receive:=booking[i].ReceiveBy
				counselor:=CheckCounSelorOnline()
				support:=counselor.PhoneSupport
				if(support==""){
					support=counselor.SipSupport
				}
				if(receive==""){
					userDto:=dao.GetUserDto(booking[i].UserClient)
					receive=userDto.SipAccount
					if(receive==""){
						receive=userDto.PhoneNumber
					}
				}
				data, _ :=http.Get("https://verdigris-puffin-4907.twil.io/useURL?To="+receive+"&Support"+support)
				//Delete data from table booking and insert data to booking_last :
				println(data)
				dao.InsertBookingLastDao(booking[i])
				dao.DeleteBookingDao(booking[i].K)
			}
		}
		fmt.Println(booking)
	}
}

func CheckCounSelorOnline()dto.CounselorOnline{
	counselor :=dao.CheckCounSelorOnlineDao()
	i:=rand.Intn(len(counselor)-1)
	dao.SetStatusCounSelor(counselor[i],"Busy");
	return counselor[i]
}