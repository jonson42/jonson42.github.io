package dao

import (
	"google.golang.org/api/iterator"
	"cloud.google.com/go/datastore"
	"context"
	"log"
	"app/models/dto"
	"app/models/form"
	"os"
	"time"
)

func CheckUserDao(username string, pass string)dto.MUser{
	ctx := context.Background()
	// Set your Google Cloud Platform project ID.
	client:=ConnectDataStore()
	userLogin:=dto.MUserAuth{}
	queryLogin:=datastore.NewQuery("m_user_auth").Namespace("TwilioAiko").Filter("id =",username).Filter("pass=",pass)
	it:=client.Run(ctx,queryLogin)
	for {
		_, err := it.Next(&userLogin)
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Error fetching next task: %v", err)
		}
	}

	queryuser:=datastore.NewQuery("m_user").Namespace("TwilioAiko").Filter("user_id =",userLogin.UserId)
	ituser:=client.Run(ctx,queryuser)
	var muser dto.MUser
	for {
		_, err := ituser.Next(&muser)
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Error fetching next task: %v", err)
		}
	}
	return muser
}

func GetUserDto(userId int)dto.MUser{
	userDto:=[]dto.MUser{}
	client:=ConnectDataStore()
	queryLogin:=datastore.NewQuery("m_user").Namespace("TwilioAiko").Filter("user_id=",userId)
	if _,err:=client.GetAll(context.Background(),queryLogin,&userDto);err!=nil{
		log.Printf("Don't find data")
	}
	return userDto[0]
}

func EditProfile(userForm form.EditProfileForm)error{
	client:=ConnectDataStore()
	tx,err:=client.NewTransaction(context.Background())
	if err!=nil{
		log.Fatalf("client.NewTransaction: %v",err)
	}
	user:= dto.MUser{}
	userKey := datastore.Key{}
	userKey.Namespace = os.Getenv("namespace")
	userKey.Kind = "m_user"
	userQuery:=datastore.NewQuery("m_user").Namespace("TwilioAiko").Filter("user_id	=",userForm.UserId)
	it:=client.Run(context.Background(),userQuery)
	for {
		_, err := it.Next(&user)
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Error fetching next task: %v", err)
		}
	}

	user.PhoneNumber = userForm.PhoneNumber
	user.SipAccount = userForm.SipAccount
	user.UserFirstName = userForm.UserFirstName
	user.UserLastName = userForm.UserLastName
	user.Amount = userForm.Amount
	user.Address = userForm.Address
	user.CountryCode = userForm.CountryCode
	user.Age = userForm.Age
	user.UpdatedAt = time.Now()
	if _,err:=tx.Put(&userKey,&user);err!=nil{
		log.Fatalf("tx.Put:%v",err)
	}
	if _,err:=tx.Commit();err!=nil{
		log.Printf("tx.Commit:%v",err)
	}
	return err
}