package dao

import (
	"context"
	"cloud.google.com/go/datastore"
	"os"
	"reflect"
)

func ConnectDataStore()*datastore.Client{
	ctx := context.Background()
	// Set your Google Cloud Platform project ID.
	projectID := os.Getenv("projectId")
	client, _ := datastore.NewClient(ctx, projectID)
	return client
}

func InputData(kind string,object reflect.Type) error{
	ctx:=context.Background()
	taskKey := datastore.Key{}
	taskKey.Namespace = os.Getenv("projectId")
	taskKey.Kind = kind
	client:=ConnectDataStore()
	 _,err := client.Put(ctx, &taskKey, &object)
	return err
}
