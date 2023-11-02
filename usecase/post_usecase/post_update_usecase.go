package post_usecase

import (
	"hackathon/dao"
	"hackathon/model"
	"log"
)

// UpdatePost update a post
func UpdatePost(p model.Post) error {
	bytes, err := dao.UpdatePost(p)
	log.Printf("bytes: %v\n", bytes)
	return err
}
