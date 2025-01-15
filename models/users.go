package models

type User struct {
	Id        string `json:"id"`
	Ism       string `json:"ism"`
	Familya   string `json:"familya"`
	Username  string `json:"username"`
	Bio       string `json:"bio"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UserCreate struct {
	Id        string `json:"id"`
	Ism       string `json:"ism"`
	Familya   string `json:"familya"`
	Username  string `json:"username"`
	Bio       string `json:"bio"`
	Email     string `json:"email"`
	UpdatedAt string `json:"updated_at"`
}

type UserUpdate struct {
	Id        string `json:"id"`
	Ism       string `json:"ism"`
	Familya   string `json:"familya"`
	Username  string `json:"username"`
	Bio       string `json:"bio"`
	Email     string `json:"email"`
	UpdatedAt string `json:"updated_at"`
}
type UserGetListRequest struct {
	Page  int64 `json:"page"`
	Limit int64 `json:"limit"`
}

type UserDelete struct {
	Id string `json:"id"`
}

type UserGetListResponse struct {
	Count int64   `json:"count"`
	Users []*User `json:"users"`
}
