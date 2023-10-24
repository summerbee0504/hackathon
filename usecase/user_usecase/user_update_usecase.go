package user_usecase

import (
	"hackathon/dao"
	"hackathon/model"
)

// GetUpdateUser update a user
func GetUpdateUser(u model.UpdateUserDetails) error {

	if err := dao.UpdateUser(u); err != nil {
		return err
	}
	return nil
}
