package model

type Post struct {
	Id           string `json:"id"`
	UserId       string `json:"user_id"`
	CategoryId   int    `json:"category_id"`
	Title        string `json:"title"`
	Url          string `json:"url"`
	Content      string `json:"content"`
	CurriculumId string `json:"curriculum_id"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

type GetPost struct {
	Id           string `json:"id"`
	Category     string `json:"category"`
	User         string `json:"user"`
	UserId       string `json:"user_id"`
	UserImage    string `json:"user_image"`
	Title        string `json:"title"`
	Url          string `json:"url"`
	Content      string `json:"content"`
	Curriculum   string `json:"curriculum"`
	CurriculumId string `json:"curriculum_id"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	CommentCount int    `json:"comment_count"`
	LikeCount    int    `json:"like_count"`
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
	UserId    string `json:"user_id"`
	UserImage string `json:"user_image"`
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
