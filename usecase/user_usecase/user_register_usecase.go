package user_usecase

import (
	"hackathon/dao"
	"hackathon/model"
)

func GetRegisterUser(u model.User) error {

	if err := dao.RegisterUser(u); err != nil {
		return err
	}

	return nil
}
