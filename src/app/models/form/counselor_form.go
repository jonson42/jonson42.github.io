package form

type CounselorForm struct {
	CounselorId int `datastore:"counselor_id" form:"counselorId"`
	SipSupport string `datastore:"sip_support" form:"sipSupport"`
	PhoneSupport string `datastore:"phone_support" form:"phoneSupport"`
}
