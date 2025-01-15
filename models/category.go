package models

type Category struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Index       int    `json:"index"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type CreateCategory struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Index       int    `json:"index"`
	Title       string `json:"title"`
	Description string `json:"description"`
	UpdatedAt   string `json:"updated_at"`
}

type UptadeCategory struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Index       int    `json:"index"`
	Title       string `json:"title"`
	Description string `json:"description"`
	UpdatedAt   string `json:"updated_at"`
}

type DeleteCategory struct {
	Id string `json:"id"`
}
