package post_usecase

import (
	"hackathon/dao"
	"hackathon/model"
	"log"
)

// DeletePost delete a post
func DeletePost(i model.SearchById) error {
	id := i.Id
	bytes, err := dao.DeletePost(id)
	if err != nil {
		return err
	}
	log.Printf("bytes: %v\n", bytes)
	return nil
}
