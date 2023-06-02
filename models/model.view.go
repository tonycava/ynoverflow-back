package models

type View struct {
	Base
	UserId string `json:"userId"`
	PostId string `json:"postId"`
}
