package user_usecase

import (
	"hackathon/dao"
	"hackathon/model"
)

func ShowUserList() (bytes []byte, err error) {
	bytes, err = dao.ShowUserList()
	return bytes, err
}

func GetShowUserDetail(i model.SearchById) (bytes []byte, err error) {
	id := i.Id
	bytes, err = dao.ShowUserDetail(id)
	return bytes, err
}
