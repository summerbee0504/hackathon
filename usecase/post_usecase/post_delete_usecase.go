package post_usecase

import "hackathon/dao"

// DeletePost delete a post
func DeletePost(id string) error {
	err := dao.DeletePost(id)
	return err
}
