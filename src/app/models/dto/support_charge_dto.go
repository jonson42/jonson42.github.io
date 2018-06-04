package dto

type SupportCharge struct {
	Id        int    `datastore:"id"`
	Amount    int    `datastore:"amount"`
	CreatedAt string `datastore:"created_at"`
	UpdatedAt string `datastore:"updated_at"`
}