package models


type Comments struct {
	Id string `json:"id"`
	UserId string `json:"user_id"`
	PostId string `json:"post_id"`
	Comment string `json:"comment"`
	CreatedAt string `json:"created_at"`
}

type CreateComment struct {
	Id string `json:"id"`
	UserId string `json:"user_id"`
	PostId string `json:"post_id"`
	Comment string `json:"comment"`
	CreatedAt string `json:"created_at"`
}
