package post_usecase

import (
	"github.com/oklog/ulid/v2"
	"hackathon/dao"
	"hackathon/model"
)

// LikePost like a post
func LikePost(l model.Like) (bytes []byte, err error) {

	ulidId := ulid.Make()
	l.Id = ulidId.String()

	bytes, err = dao.LikePost(l)
	return bytes, err
}

// UnlikePost unlike a post
func UnlikePost(l model.Unlike) (bytes []byte, err error) {
	bytes, err = dao.UnlikePost(l)
	return bytes, err
}
