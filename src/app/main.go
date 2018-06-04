package main

import(
	"github.com/gin-gonic/gin"
	"app/controller"
	"google.golang.org/appengine"
)

func main(){
	appengine.Main()
}

func Main(){

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	//api
	r.POST("/login",controller.Login)
	r.POST("/register",controller.Register)
	r.GET("/logout",controller.Logout)
	r.POST("/booking",controller.Booking)
	r.GET("/call-process",controller.CallProcess)
	r.GET("/list-call/:userId",controller.GetListCall)
	r.GET("/profile/:userId",controller.GetProfile)
	r.POST("/edit-booking",controller.EditBooking)
	r.POST("/edit-profile",controller.EditProfile)
	r.POST("/admin/set-counselor",controller.SetCounselor)
	//webpage
	r.GET("/main",controller.MainPage)
	r.GET("/",controller.IndexPage)
	r.GET("/register",controller.RegisterPage)
	r.GET("/profile",controller.ProfilePage)
	r.Run(":8080")
}


