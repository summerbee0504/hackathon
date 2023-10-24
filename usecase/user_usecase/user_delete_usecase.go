package user_usecase

import (
	"hackathon/dao"
)

// GetDeleteUser delete a user
func GetDeleteUser(id string) error {
	return dao.DeleteUser(id)
}
