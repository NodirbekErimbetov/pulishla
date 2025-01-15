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
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
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
	UpdatedAt   string `json:"updated_at"`
}
type UpdateView struct {
	Id          string `json:"id"`
	Count       int    `json:"count"`
	Like        int64  `json:"like"`
	Complated   int64  `json:"complated"`
	WatchedTime int64  `json:"watched_time"`
	Star        int64  `json:"star"`
	UpdatedAt   string `json:"updated_at"`
}

type PrimaryKeyView struct {
	Id string `json:"id"`
}

type GetListViewResponse struct {
	Count int64 `json:"count"`
	View []*Views `json:"view"`
}
