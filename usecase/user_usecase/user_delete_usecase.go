package user_usecase

import (
	"hackathon/dao"
	"hackathon/model"
)

// GetDeleteUser delete a user
func GetDeleteUser(i model.SearchById) (bytes []byte, err error) {
	id := i.Id
	bytes, err = dao.DeleteUser(id)
	return bytes, err
}
