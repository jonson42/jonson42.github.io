package controller

import (
	"github.com/gin-gonic/gin"
	"app/models/form"
	"net/http"
	"app/models/dao"
)
func SetCounselor(c *gin.Context){
	counselor:=form.CounselorForm{}
	if err:=c.Bind(&counselor);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err})
	}
	result:=dao.SetCounSelorDao(counselor)
	if(result==nil){
		c.JSON(http.StatusOK,gin.H{"result":"true"})
	}else{
		c.JSON(http.StatusBadRequest,gin.H{"result":"false"})
	}
}
