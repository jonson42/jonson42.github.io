package controller

import(
	"github.com/gin-gonic/gin"
	"app/models/dao"
	"strconv"
	"net/http"
)

func GetListCall(c *gin.Context){
	userId,_:=strconv.Atoi(c.Param("userId"))
	result:=dao.GetListDao(userId)
	if(len(result)!=0){
		c.JSON(http.StatusOK,gin.H{"datas":result})
	}else{
		c.JSON(http.StatusNotFound,gin.H{"error":"Data not found"})
	}
}
