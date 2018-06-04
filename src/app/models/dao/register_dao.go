package dao

import (
	"github.com/gin-gonic/gin"
	"app/models/dto"
	"app/models/form"
	"os"
	"context"
	"cloud.google.com/go/datastore"
	"google.golang.org/api/iterator"
	"log"
	"time"
)

func RegisterDao(c *gin.Context,form form.RegiserForm)error{
	user:=dto.MUser{}
	userAuth:=dto.MUserAuth{}
	userId:=GetCurrentUserNumber()+1
	user.UserFirstName=form.FirstName
	user.UserLastName=form.LastName
	user.PhoneNumber = form.PhoneNumber
	user.SipAccount = form.SipAccount
	user.CountryCode = form.CountryCode
	user.UserId=userId
	user.UpdatedAt = time.Now()
	user.CreatedAt = time.Now()

	userAuth.UserId = userId
	userAuth.Id = form.UserName
	userAuth.Pass = form.Password
	userAuth.UpdatedAt = time.Now()
	userAuth.CreatedAt = time.Now()
	// update data to datastore
	ctx:=context.Background()
	taskUser := datastore.Key{}
	taskUser.Namespace = os.Getenv("namespace")
	taskUser.Kind = "m_user"
	taskUserAuth := datastore.Key{}
	taskUserAuth.Namespace = os.Getenv("namespace")
	taskUserAuth.Kind = "m_user_auth"
	client:=ConnectDataStore()
	// update table m_user
	_,err := client.Put(ctx, &taskUser, &user)
    if(err!=nil){
    	return err
	}
	// update table m_user_auth
	_,err1:=client.Put(ctx,&taskUserAuth,&userAuth)
	return err1
}

func GetCurrentUserNumber()int{
	dataNeed:=0
	ctx:=context.Background()
	projectID := os.Getenv("projectId")
	client, _ := datastore.NewClient(ctx, projectID)
	userIdCurrent:=dto.AutoIdStart{}
	queryLogin:=datastore.NewQuery("auto_id_start").Namespace("TwilioAiko")
	it:=client.Run(ctx,queryLogin)
	for {
		_, err := it.Next(&userIdCurrent)
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Error fetching next task: %v", err)
		}
	}
	//update update m_user_auth
	tx,err:=client.NewTransaction(ctx)
	if err!=nil{
		log.Fatalf("client.NewTransaction: %v",err)
	}

	dataNeed = userIdCurrent.UserIdCurrent

	task := datastore.Key{}
	task.Namespace = os.Getenv("namespace")
	task.Kind = "auto_id_start"
	if err := tx.Get(userIdCurrent.K, &userIdCurrent); err != nil {
		log.Fatalf("tx.Get: %v", err)
	}
	userIdCurrent.UserIdCurrent = userIdCurrent.UserIdCurrent+1
	userIdCurrent.UpdateAt = time.Now()
	if _,err:=tx.Put(userIdCurrent.K,&userIdCurrent);err!=nil{
		log.Fatalf("tx.Put:%v",err)
	}
	if _,err:=tx.Commit();err!=nil{
		log.Printf("tx.Commit:%v",err)
	}

	return dataNeed

}