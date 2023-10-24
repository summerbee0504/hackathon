package post_usecase

import (
	"github.com/oklog/ulid/v2"
	"hackathon/dao"
	"hackathon/model"
)

// MakeComment make a new comment on post
func MakeComment(c model.Comment) error {
	ulidId := ulid.Make()
	c.Id = ulidId.String()

	err := dao.CommentPost(c)
	return err
}

func GetComments(id string) (bytes []byte, err error) {
	bytes, err = dao.GetAllCommentsByPost(id)
	return bytes, err
}

// DeleteComment delete a comment on post
func DeleteComment(id string) error {
	err := dao.DeleteComment(id)
	return err
}

// UpdateComment update a comment on post
func UpdateComment(c model.Comment) error {
	err := dao.UpdateComment(c)
	return err
}
