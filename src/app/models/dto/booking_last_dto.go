package dto

import "time"

type BookingLast struct {
	Id int `datastore:"id" json:"id"`
	Status string `datastore:"status" json:"status"`
	CounselorSupport string `datastore:"counselor_support" json:"counselorSupport"`
	UserClient int `datastore:"user_client" json:"userClient"`
	DateTime string `datastore:"data_time" json:"dataTime"`
	TimeCall string `datastore:"time_call" json:"timeCall"`
	MoneyPaid string `datastore:"money_paid" json:"moneyPaid"`
	Note string `datastore:"note" json:"note"`
	UpdatedAt time.Time `datastore:"updated_at" json:"updatedAt"`
	CreatedAt time.Time `datastore:"created_at" json:"createdAt"`

}