package dto

import (
	 "cloud.google.com/go/datastore"
	"time"
)

type AutoIdStart struct {
	K *datastore.Key `datastore:"__key__"`
	UserIdCurrent int `datastore:"user_id_current" form:""`
	UpdateAt time.Time `datastore:"update_at"`
	CreateAt time.Time `datastore:"create_at"`
}