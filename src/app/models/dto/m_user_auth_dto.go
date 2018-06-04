package dto

import (
	"time"

	  )

type MUserAuth struct {
	UserId int `datastore:"user_id"`
	Id string `datastore:"id"`
	Pass string `datastore:"pass"`
	UpdatedAt time.Time `datastore:"updated_at"`
	CreatedAt time.Time `datastore:"created_at"`
}
