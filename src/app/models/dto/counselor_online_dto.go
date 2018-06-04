package dto

import (
	"cloud.google.com/go/datastore"
	"time"
)

type CounselorOnline struct {
	Id *datastore.Key `datastore:"__key__"`
	CounselorId int `datastore:"counselor_id"`
	SipSupport string `datastore:"sip_support"`
	PhoneSupport string `datastore:"phone_support"`
	Status string `datastore:"status"`
	UpdatedAt time.Time `datastore:"updated_at"`
	CreatedAt time.Time `datastore:"created_at"`
}