package user_usecase

import (
	"hackathon/dao"
	"hackathon/model"
	"log"
)

// GetUpdateUser update a user
func GetUpdateUser(u model.UpdateUserDetails) error {

	bytes, err := dao.UpdateUser(u)
	if err != nil {
		return err
	}
	log.Printf("%v\n", bytes)
	return nil
}
