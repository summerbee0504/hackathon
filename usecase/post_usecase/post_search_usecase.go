package post_usecase

import "hackathon/dao"

// GetPost get a post
func GetPost(id string) (bytes []byte, err error) {
	bytes, err = dao.GetPost(id)
	return bytes, err
}

// GetLikeCount get like count of a post
func GetLikeCount(id string) (bytes []byte, err error) {
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
func GetAllPostsByUser(id string) (bytes []byte, err error) {
	bytes, err = dao.GetAllPostsByUser(id)
	return bytes, err
}

// GetAllPostsByTag get all posts by tag
func GetAllPostsByTag(id string) (bytes []byte, err error) {
	bytes, err = dao.GetAllPostsByTag(id)
	return bytes, err
}

// GetAllPostsByCategory get all posts by category
func GetAllPostsByCategory(id int) (bytes []byte, err error) {
	bytes, err = dao.GetAllPostsByCategory(id)
	return bytes, err
}

// GetAllPostsByCurriculum get all posts by curriculum
func GetAllPostsByCurriculum(id int) (bytes []byte, err error) {
	bytes, err = dao.GetAllPostsByCurriculum(id)
	return bytes, err
}

// GetAllPostsByDate get all posts by date
func GetAllPostsByDate() (bytes []byte, err error) {
	bytes, err = dao.GetAllPostsByDate()
	return bytes, err
}
