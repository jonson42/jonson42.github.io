package form

import "time"

type EditProfileForm struct {
	UserId int `form:"userId"`
	PhoneNumber string `form:"phoneNumber"`
	SipAccount string `form:"sipAccount"`
	UserFirstName string `form:"userFirstName"`
	UserLastName string `form:"userLastName"`
	Amount string `form:"amount"`
	Address string `form:"Address"`
	CountryCode string `form:"countryCode"`
	Age int `form:"age"`
	UserLevl int `form:"userLevl"`
	UpdatedAt time.Time `form:"updatedAt"`
	CreatedAt time.Time `form:"createdAt"`
}