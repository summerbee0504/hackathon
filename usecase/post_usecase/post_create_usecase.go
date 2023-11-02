package post_usecase

import (
	"github.com/oklog/ulid/v2"
	"hackathon/dao"
	"hackathon/model"
	"log"
)

// MakePost make a new post
func MakePost(p model.Post) error {

	ulidId := ulid.Make()
	p.Id = ulidId.String()

	bytes, err := dao.MakePost(p)

	log.Printf("bytes: %v\n", bytes)

	return err
}
