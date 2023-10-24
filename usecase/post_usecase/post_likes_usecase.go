package post_usecase

import (
	"github.com/oklog/ulid/v2"
	"hackathon/dao"
	"hackathon/model"
)

// LikePost like a post
func LikePost(l model.Like) error {

	ulidId := ulid.Make()
	l.Id = ulidId.String()

	err := dao.LikePost(l)
	return err
}

// UnlikePost unlike a post
func UnlikePost(id string) error {
	err := dao.UnlikePost(id)
	return err
}

func GetLikedPostByUser(id string) (bytes []byte, err error) {
	bytes, err = dao.GetLikedPosts(id)
	return bytes, err
}
