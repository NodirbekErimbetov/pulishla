package models

type Views struct {
	Id          string `json:"id"`
	PostId      string `json:"post_id"`
	UserId      string `json:"user_id"`
	Count       int    `json:"count"`
	Like        int64  `json:"like"`
	Complated   int64  `json:"complated"`
	WatchedTime int64  `json:"watched_time"`
	Star        int64  `json:"star"`
}

type CreateView struct {
	Id          string `json:"id"`
	PostId      string `json:"post_id"`
	UserId      string `json:"user_id"`
	Count       int    `json:"count"`
	Like        int64  `json:"like"`
	Complated   int64  `json:"complated"`
	WatchedTime int64  `json:"watched_time"`
	Star        int64  `json:"star"`
}

