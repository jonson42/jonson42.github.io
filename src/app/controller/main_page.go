package controller

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"strconv"
	_"html/template"
	"net/http"
)

func MainPage(c *gin.Context){
	check:=CheckSessionLogin(c)
	if check {
		userId:=GetUserId(c.Request)
		fmt.Println("UserId : "+strconv.Itoa(userId))
		c.HTML(http.StatusOK,"main.html",nil)
		//templatePage := template.Must(template.ParseFiles("app/templates/main.html"))
		//templatePage.Execute(c.Writer, nil)
	}else{
		c.Redirect(302, "/")
	}

}