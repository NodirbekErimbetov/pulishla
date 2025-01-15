package models

type Comment struct {
	Id        string `json:"id"`
	UserId    string `json:"user_id"`
	PostId    string `json:"post_id"`
	Comment   string `json:"comment"`
	CreatedAt string `json:"created_at"`
}

type CreateComment struct {
	Id        string `json:"id"`
	UserId    string `json:"user_id"`
	PostId    string `json:"post_id"`
	Comment   string `json:"comment"`
	CreatedAt string `json:"created_at"`
}

type DeleteComment struct {
	Id string `json:"id"`
}

type GetListCommentResponse struct {
	Count int64 `json:"count"`
	Comments []*Comment `json:"comments"`
}
