package controller
import (
	"github.com/gin-gonic/gin"
	_"html/template"
	"net/http"
)

func IndexPage(c *gin.Context){
	check:=CheckSessionLogin(c)
	if check {
		c.HTML(http.StatusOK,"main.html",nil)
		//templatePage := template.Must(template.ParseFiles("app/templates/main.html"))
		//templatePage.Execute(c.Writer, nil)
	}else{
		c.HTML(http.StatusOK,"index.html",nil)
		//templatePage := template.Must(template.ParseFiles("app/templates/index.html"))
		//templatePage.Execute(c.Writer, nil)
	}
}
