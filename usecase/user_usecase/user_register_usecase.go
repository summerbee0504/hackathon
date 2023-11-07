package user_usecase

import (
	"hackathon/dao"
	"hackathon/model"
)

func GetRegisterUser(u model.User) (bytes []byte, err error) {

	bytes, err = dao.RegisterUser(u)
	return bytes, err
}
