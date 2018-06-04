package dto

import "time"

type MUser struct {
	UserId int `datastore:"user_id" json:"userId"`
	PhoneNumber string `datastore:"phone_number" json:"phoneNumber"`
	SipAccount string `datastore:"sip_account" json:"sipAccount"`
	UserFirstName string `datastore:"user_first_name" json:"userFirstName"`
	UserLastName string `datastore:"user_last_name" json:"userLastName"`
	Amount string `datastore:"amount" json:"amount" json:"amount"`
	Address string `datastore:"adresss" json:"address"`
	CountryCode string `datastore:"country_code" json:"countryCode"`
	Age int `datastore:"age" json:"age"`
	UserLevl int `datastore:"user_levl" json:"userLevl"`
	UpdatedAt time.Time `datastore:"updated_at" json:"updatedAt"`
	CreatedAt time.Time `datastore:"created_at" json:"createdAt"`
}