package post_usecase

import (
	"hackathon/dao"
	"hackathon/model"
)

// UpdatePost update a post
func UpdatePost(p model.Post) error {
	err := dao.UpdatePost(p)
	return err
}
