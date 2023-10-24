package post_usecase

import (
	"github.com/oklog/ulid/v2"
	"hackathon/dao"
	"hackathon/model"
)

// MakePost make a new post
func MakePost(p model.Post) error {

	ulidId := ulid.Make()
	p.Id = ulidId.String()

	err := dao.MakePost(p)
	return err
}
