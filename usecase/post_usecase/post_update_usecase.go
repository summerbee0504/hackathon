package post_usecase

import (
	"hackathon/dao"
	"hackathon/model"
)

// UpdatePost update a post
func UpdatePost(p model.Post) (bytes []byte, err error) {
	bytes, err = dao.UpdatePost(p)
	return bytes, err
}
