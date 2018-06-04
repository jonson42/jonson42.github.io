package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"app/models/dao"
	"app/models/form"
)

func Login(c *gin.Context){
	/*log.Println("login")
	formPost:=form.MUserAuthForm{}
	if err:=c.Bind(&formPost);err !=nil {
		c.JSON(http.StatusConflict,gin.H{"error":err.Error()})
		return
	}
	user:=dao.CheckUserDao(formPost.UserName,formPost.Password)
	if user.UserId!=0{
		SetSession(strconv.Itoa(user.UserId),c.Writer)
		c.Redirect(http.StatusMovedPermanently, "/main")
	}else {
		c.JSON(http.StatusNotFound,"Page not found")
	}*/
	c.JSON(http.StatusNotFound,"Page not found")
}

func Logout(c *gin.Context){
	ClearSession(c.Writer)
	c.Redirect(http.StatusFound, "/")
}

func Register(c *gin.Context){
	registerForm:=form.RegiserForm{}
	if err:=c.Bind(&registerForm);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err})
		return
	}
	err:=dao.RegisterDao(c,registerForm)
	if(err==nil){
		c.Redirect(http.StatusMovedPermanently,"/main")
	}else{
		c.JSON(http.StatusNotFound,gin.H{"error":err})
	}
}

func RegisterPage(c *gin.Context){
	if !CheckSessionLogin(c){
		c.HTML(http.StatusOK,"register.html",nil)
		//templatePage := template.Must(template.ParseFiles("src/app/templates/posts/register.html"))
		//templatePage.Execute(c.Writer, nil)
	}
}



