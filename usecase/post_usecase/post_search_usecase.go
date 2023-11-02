package post_usecase

import (
	"hackathon/dao"
	"hackathon/model"
	"strconv"
)

// GetPost get a post
func GetPost(i model.SearchById) (bytes []byte, err error) {
	id := i.Id
	bytes, err = dao.GetPost(id)
	return bytes, err
}

// GetLikeCount get like count of a post
func GetLikeCount(i model.SearchById) (bytes []byte, err error) {
	id := i.Id
	bytes, err = dao.GetLikeCount(id)
	return bytes, err
}

func GetAllCurriculums() (bytes []byte, err error) {
	bytes, err = dao.GetAllCurriculums()
	return bytes, err
}

func GetAllTags() (bytes []byte, err error) {
	bytes, err = dao.GetAllTags()
	return bytes, err
}

// GetAllPostsByUser get all posts by user
func GetAllPostsByUser(i model.SearchById) (bytes []byte, err error) {
	id := i.Id
	bytes, err = dao.GetAllPostsByUser(id)
	return bytes, err
}

// GetAllPostsByTag get all posts by tag
func GetAllPostsByTag(i model.SearchById) (bytes []byte, err error) {
	id := i.Id
	bytes, err = dao.GetAllPostsByTag(id)
	return bytes, err
}

// GetAllPostsByCategory get all posts by category
func GetAllPostsByCategory(i model.SearchById) (bytes []byte, err error) {
	id, errMsg := strconv.Atoi(i.Id)
	if err != nil {
		return nil, errMsg
	}
	bytes, err = dao.GetAllPostsByCategory(id)
	return bytes, err
}

// GetAllPostsByCurriculum get all posts by curriculum
func GetAllPostsByCurriculum(i model.SearchById) (bytes []byte, err error) {
	id, errMsg := strconv.Atoi(i.Id)
	if err != nil {
		return nil, errMsg
	}
	bytes, err = dao.GetAllPostsByCurriculum(id)
	return bytes, err
}

// GetAllPostsByDate get all posts by date
func GetAllPostsByDate() (bytes []byte, err error) {
	bytes, err = dao.GetAllPostsByDate()
	return bytes, err
}
