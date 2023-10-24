package user_usecase

import (
	"hackathon/dao"
)

func ShowUserList() (bytes []byte, err error) {
	bytes, err = dao.ShowUserList()
	return bytes, err
}

func ShowUserDetail(id string) (bytes []byte, err error) {
	bytes, err = dao.ShowUserDetail(id)
	return bytes, err
}
