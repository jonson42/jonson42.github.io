package controller
import (
	"net/http"
	"github.com/gorilla/securecookie"
	"github.com/gin-gonic/gin"
	"strconv"
	"google.golang.org/api/iterator"
	"log"
	"cloud.google.com/go/datastore"
	"app/models/dao"
	"app/models/dto"
	"golang.org/x/net/context"
)

var cookieHandler = securecookie.New(
	   securecookie.GenerateRandomKey(64),
       securecookie.GenerateRandomKey(32))
func SetSession(userId string, response http.ResponseWriter) {
	value := map[string]string{
		"userId": userId,
	}
	if encoded, err := cookieHandler.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:  "session",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(response, cookie)
	}
}

func ClearSession(response http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)
}

func GetUserId(request *http.Request) (userId int) {
	if cookie, err := request.Cookie("session"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {
			userId,_ =strconv.Atoi(cookieValue["userId"])
		}
	}
	return userId
}

func GetUserDtoFromSession(request *http.Request)dto.MUser{
	userDto:=dto.MUser{}
	userId:=GetUserId(request)
	queryLogin:=datastore.NewQuery("m_user").Namespace("TwilioAiko").Filter("user_id=",userId)
	client:=dao.ConnectDataStore()
	it:=client.Run(context.Background(),queryLogin)
	for {
		_, err := it.Next(&userDto)
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Error fetching next task: %v", err)
		}
	}
	return userDto
}

func CheckSessionLogin(c *gin.Context)bool{
	userName:=GetUserId(c.Request)
	if userName==0{
		return false
	}
	return true
}