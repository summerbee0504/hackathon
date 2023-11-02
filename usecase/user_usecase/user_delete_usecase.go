package user_usecase

import (
	"hackathon/dao"
	"hackathon/model"
	"log"
)

// GetDeleteUser delete a user
func GetDeleteUser(i model.SearchById) error {
	id := i.Id
	bytes, err := dao.DeleteUser(id)
	if err != nil {
		return err
	}
	log.Printf("%v\n", bytes)
	return nil
}
