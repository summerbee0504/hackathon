package user_usecase

import (
	"github.com/oklog/ulid/v2"
	"hackathon/dao"
	"hackathon/model"
)

func GetRegisterUser(u model.User) error {

	ulidId := ulid.Make()
	u.Id = ulidId.String()

	if err := dao.RegisterUser(u); err != nil {
		return err
	}

	return nil
}
