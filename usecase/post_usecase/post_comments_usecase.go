package post_usecase

import (
	"github.com/oklog/ulid/v2"
	"hackathon/dao"
	"hackathon/model"
)

// MakeComment make a new comment on post
func MakeComment(c model.Comment) (bytes []byte, err error) {
	ulidId := ulid.Make()
	c.Id = ulidId.String()

	bytes, err = dao.CommentPost(c)
	return bytes, err
}

func GetComments(i model.SearchById) (bytes []byte, err error) {
	id := i.Id
	bytes, err = dao.GetAllCommentsByPost(id)
	return bytes, err
}

// DeleteComment delete a comment on post
func DeleteComment(i model.SearchById) (bytes []byte, err error) {
	id := i.Id
	bytes, err = dao.DeleteComment(id)
	return bytes, err
}

// UpdateComment update a comment on post
func UpdateComment(c model.Comment) (bytes []byte, err error) {
	bytes, err = dao.UpdateComment(c)
	return bytes, err
}
