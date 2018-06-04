package form

type RegiserForm struct {
	FirstName string `form:"firstName"`
	LastName string `form:"lastName"`
	UserName string	`form:"userName"`
	Password string `form:"password"`
	Email string `form:"email"`
	PhoneNumber string `form:"phoneNumber"`
	SipAccount string `form:"sipAccount"`
	CountryCode string `form:"countryCode"`
}
