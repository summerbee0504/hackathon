package post_usecase

import (
	"hackathon/dao"
	"hackathon/model"
)

// DeletePost delete a post
func DeletePost(i model.SearchById) (bytes []byte, err error) {
	id := i.Id
	bytes, err = dao.DeletePost(id)
	return bytes, err
}
