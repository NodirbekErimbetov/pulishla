package models

type Post struct {
	Id           string  `json:"id"`
	CategoryId   string  `json:"category_id"`
	Body         string  `json:"body"`
	Title        string  `json:"title"`
	DuratingTime int64   `json:"duration_time"`
	Stars        float64 `json:"stars"`
	About        string  `json:"about"`
	Alt          string  `json:"alt"`
	Meta         string  `json:"meta"`
	VideoUrl     string  `json:"video_url"`
	Level        int     `json:"level"`
	ViewsCount   int     `json:"views_count"`
	StarCount    int     `json:"star_count"`
	CreatedAt    string  `json:"created_at"`
	UptadetAt    string  `json:"uptadet_at"`
}

type CreatePost struct {
	Id           string  `json:"id"`
	CategoryId   string  `json:"category_id"`
	Body         string  `json:"body"`
	Title        string  `json:"title"`
	DuratingTime int64   `json:"duration_time"`
	Stars        float64 `json:"stars"`
	About        string  `json:"about"`
	Alt          string  `json:"alt"`
	Meta         string  `json:"meta"`
	VideoUrl     string  `json:"video_url"`
	Level        int     `json:"level"`
	ViewsCount   int     `json:"views_count"`
	StarCount    int     `json:"star_count"`
	UptadetAt    string  `json:"uptadet_at"`
}

type UpdatePost struct {
	Id           string `json:"id"`
	CategoryId   string `json:"category_id"`
	Body         string `json:"body"`
	Title        string `json:"title"`
	DuratingTime int64   `json:"duration_time"`
	Stars        float64 `json:"stars"`
	About        string  `json:"about"`
	Alt          string  `json:"alt"`
	Meta         string  `json:"meta"`
	VideoUrl     string  `json:"video_url"`
	Level        int     `json:"level"`
	ViewsCount   int     `json:"views_count"`
	StarCount    int     `json:"star_count"`
	UptadetAt    string `json:"uptadet_at"`
}

type DeletePost struct {
	Id string `json:"id"`
}

type GetListPostRequest struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

type GetListPostResponse struct {
	Count int     `json:"count"`
	Post  []*Post `json:"post"`
}
