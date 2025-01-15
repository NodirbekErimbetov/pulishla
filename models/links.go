package models

type Links struct{
	Id string `json:"id"`
	PostId string `json:"post_id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Links string `json:"links"`
	Referal string `json:"referal"`
}

type CreateLinks struct {
	Id string `json:"id"`
	PostId string `json:"post_id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Links string `json:"links"`
	Referal string `json:"referal"`
	CreatedAt string `json:"created_at"`
}