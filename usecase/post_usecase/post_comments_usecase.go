package post_usecase

import (
	"github.com/oklog/ulid/v2"
	"hackathon/dao"
	"hackathon/model"
	"log"
)

// MakeComment make a new comment on post
func MakeComment(c model.Comment) error {
	ulidId := ulid.Make()
	c.Id = ulidId.String()

	bytes, err := dao.CommentPost(c)
	if err != nil {
		return err
	}
	log.Printf("%s", bytes)
	return nil
}

func GetComments(i model.SearchById) (bytes []byte, err error) {
	id := i.Id
	bytes, err = dao.GetAllCommentsByPost(id)
	return bytes, err
}

// DeleteComment delete a comment on post
func DeleteComment(i model.SearchById) error {
	id := i.Id
	bytes, err := dao.DeleteComment(id)
	if err != nil {
		return err
	}
	log.Printf("%s", bytes)
	return nil
}

// UpdateComment update a comment on post
func UpdateComment(c model.Comment) error {
	bytes, err := dao.UpdateComment(c)
	if err != nil {
		return err
	}
	log.Printf("%s", bytes)
	return nil
}
