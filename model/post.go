package model

type Post struct {
	Id           string `json:"id"`
	CategoryId   int    `json:"category_id"`
	UserId       string `json:"user_id"`
	Title        string `json:"title"`
	Content      string `json:"content"`
	CurriculumId int    `json:"curriculum_id"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

type GetPost struct {
	Id         string `json:"id"`
	Category   string `json:"category"`
	User       string `json:"user"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Curriculum string `json:"curriculum"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type SearchById struct {
	Id string `json:"id"`
}

type Like struct {
	Id      string `json:"id"`
	PostId  string `json:"post_id"`
	UserId  string `json:"user_id"`
	LikedAt string `json:"liked_at"`
}

type Unlike struct {
	PostId string `json:"post_id"`
	UserId string `json:"user_id"`
}

type Comment struct {
	Id        string `json:"id"`
	PostId    string `json:"post_id"`
	UserId    string `json:"user_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type GetComment struct {
	Id        string `json:"id"`
	PostId    string `json:"post_id"`
	User      string `json:"user"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type Curriculum struct {
	Id         string `json:"id"`
	Curriculum string `json:"curriculum"`
}

type Tag struct {
	Id  string `json:"id"`
	Tag string `json:"tag"`
}
