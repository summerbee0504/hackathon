package user_usecase

import (
	"hackathon/dao"
	"hackathon/model"
	"log"
)

func GetRegisterUser(u model.User) error {

	bytes, err := dao.RegisterUser(u)
	if err != nil {
		return err
	}
	log.Printf("%v\n", bytes)
	return nil
}
