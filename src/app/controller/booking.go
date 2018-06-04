package controller

import (
	"github.com/gin-gonic/gin"
	"app/models/form"
	"net/http"
	"app/models/dao"
)

func Booking(c *gin.Context){
  	booking:=form.BookingForm{}
  	if err:=c.Bind(&booking);err!=nil{
  		c.JSON(http.StatusBadRequest,gin.H{"Error":err.Error()})
  		return
	}
	userId:=GetUserId(c.Request)
	err:=dao.BookingDao(booking,userId)
	if err==nil{
		c.Redirect(302,"/main")
	}else{
		c.Redirect(302,"/main")
	}
}

func EditBooking(c *gin.Context){
	booking:=form.BookingForm{}
	if err:=c.Bind(&booking);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"Error":err.Error()})
		return
	}

	err:=dao.EditBookingDao(booking)
	if err==nil{
		c.Redirect(302,"/main")
	}else{
		c.Redirect(302,"/main")
	}
}
