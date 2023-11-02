package post_usecase

import (
	"github.com/oklog/ulid/v2"
	"hackathon/dao"
	"hackathon/model"
	"log"
)

// LikePost like a post
func LikePost(l model.Like) error {

	ulidId := ulid.Make()
	l.Id = ulidId.String()

	err := dao.LikePost(l)
	return err
}

// UnlikePost unlike a post
func UnlikePost(l model.Unlike) error {
	bytes, err := dao.UnlikePost(l)
	log.Printf("bytes: %v\n", bytes)
	return err
}

func GetLikedPostByUser(i model.SearchById) (bytes []byte, err error) {
	id := i.Id
	bytes, err = dao.GetLikedPosts(id)
	return bytes, err
}
