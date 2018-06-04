package dto

type MUserPictures struct {
	UserPictureId int `datastore:"user_picture_id"`
	UserId int `datastore:"user_id"`
	UserPicture string `datastore:"user_picture"`
	CreatedAt string `datastore:"created_at"`
	UpdatedAt string `datastore:"updated_at"`
}