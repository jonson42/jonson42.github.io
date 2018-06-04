package controller

import (
	"github.com/gin-gonic/gin"
	"app/models/form"
	"app/models/dao"
	"net/http"
	"strconv"
	"html/template"
)

func EditProfile(c *gin.Context){
	user:=form.EditProfileForm{}
	if err:=c.Bind(&user);err!=nil{
		c.JSON(http.StatusNotFound,gin.H{"Error":"data not found"})
		return
	}
	err:=dao.EditProfile(user)
	if(err!=nil){
		c.JSON(http.StatusInternalServerError,gin.H{"result":"error"})
	}
	c.JSON(http.StatusOK,gin.H{"result":"sucessfull"})
}

func GetProfile(c *gin.Context){
	userId,_:=strconv.Atoi(c.Param("userId"))
	user:=dao.GetUserDto(userId)
	if(user.UserId!=0){
		c.JSON(http.StatusOK,user)
	}else{
		c.JSON(http.StatusNotFound,gin.H{"error":"Data not found !"})
	}
}

func ProfilePage(c *gin.Context){
	userDto:=GetUserDtoFromSession(c.Request)
	templatePage := template.Must(template.ParseFiles("templates/profile.html"))
	templatePage.Execute(c.Writer, userDto)
}