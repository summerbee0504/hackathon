package user_usecase

import (
	"hackathon/dao"
	"hackathon/model"
)

// GetUpdateUser update a user
func GetUpdateUser(u model.UpdateUserDetails) (bytes []byte, err error) {

	bytes, err = dao.UpdateUser(u)
	return bytes, err
}
