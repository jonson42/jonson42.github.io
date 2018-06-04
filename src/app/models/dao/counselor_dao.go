package dao

import (
	"os"
	"cloud.google.com/go/datastore"
	"context"
	"log"
	"app/models/dto"
	"app/models/form"
	"time"
)

func CheckCounSelorOnlineDao()[]dto.CounselorOnline{
	projectID := os.Getenv("projectId")
	client, _ := datastore.NewClient(context.Background(), projectID)
	counselor:=[]dto.CounselorOnline{}
	queryLogin:=datastore.NewQuery("counselor_online").Namespace("TwilioAiko").Filter("status=","Active")
	if _,err:=client.GetAll(context.Background(),queryLogin,&counselor);err!=nil{
		log.Printf("Don't find data")
	}
	return counselor
}

func SetStatusCounSelor(counselor dto.CounselorOnline,status string)error{
	client:=ConnectDataStore()
	tx,err:=client.NewTransaction(context.Background())
	if err!=nil{
		log.Fatalf("client.NewTransaction: %v",err)
	}
	counselor.Status = status
	if _,err:=tx.Put(counselor.Id,&counselor);err!=nil{
		log.Fatalf("tx.Put:%v",err)
	}
	if _,err:=tx.Commit();err!=nil{
		log.Printf("tx.Commit:%v",err)
	}
	return err
}

func SetCounSelorDao(form form.CounselorForm)error{
	counselor:=dto.CounselorOnline{}
	counselor.Status = "Active"
	counselor.SipSupport =form.SipSupport
	counselor.PhoneSupport = form.PhoneSupport
	counselor.CounselorId = form.CounselorId
	counselor.UpdatedAt = time.Now()
	counselor.CreatedAt = time.Now()
	taskCounselor := datastore.Key{}
	taskCounselor.Namespace = os.Getenv("namespace")
	taskCounselor.Kind = "counselor_online"
	client:=ConnectDataStore()
	// update table counselor_online
	_,err := client.Put(context.Background(), &taskCounselor, &counselor)
	if(err!=nil){
		return err
	}
	return nil
}
