package dto

type WorkingDate struct {
	Id int `datastore:"id"`
	UserId int `datastore:"user_id"`
	StartTime string `datastore:"start_time"`
	EndTime string `datastore:"end_time"`
	Sunday string `datastore:"sunday"`
	Monday string `datastore:"monday"`
	Tuesday string `datastore:"tuesday"`
	Wednesday string `datastore:"wednesday"`
	Thursday string `datastore:"thursday"`
	Friday string `datastore:"friday"`
	Saturday string `datastore:"saturday"`
	Created_at string `datastore:"created_at"`
	Updated_at string `datastore:"updated_at"`
}