package dao

import (
	"encoding/json"
	"fmt"
	"hackathon/model"
)

func MakePost(p model.Post) (bytes []byte, err error) {
	tx, err := Db.Begin()
	if err != nil {
		return nil, fmt.Errorf("fail: Db.Begin, %v", err)
	}
	if _, err := tx.Exec("INSERT INTO posts (id, user_id, category_id, title, url, content, curriculum_id) "+
		"VALUES (?, ?, ?, ?, ?, ?, ?	)",
		p.Id, p.UserId, p.CategoryId, p.Title, p.Url, p.Content, p.CurriculumId); err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			return nil, fmt.Errorf("fail: db.Rollback, %v", rollbackErr)
		}
		return nil, fmt.Errorf("fail: db.Exec, %v", err)
	}
	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("fail: db.Commit, %v", err)
	}

	response := map[string]string{"post_id": p.Id}
	bytes, err = json.Marshal(response)
	if err != nil {
		return nil, fmt.Errorf("fail: json.Marshal, %v", err)
	}

	return bytes, nil
}

func GetPost(id string) (bytes []byte, err error) {
	rows, err := Db.Query("SELECT ca.category, u.name, u.id, u.image, p.title, p.url, p.content, cu.curriculum, cu.id, p.created_at, p.updated_at "+
		"FROM posts p "+
		"INNER JOIN categories ca ON p.category_id = ca.id "+
		"INNER JOIN users u ON p.user_id = u.id "+
		"INNER JOIN curriculums cu ON p.curriculum_id = cu.id "+
		"WHERE p.id = ?", id)
	if err != nil {
		return nil, fmt.Errorf("fail: db.Query, %v", err)
	}
	defer rows.Close()

	post := model.GetPost{}
	for rows.Next() {
		if err := rows.Scan(&post.Category, &post.User, &post.UserId, &post.UserImage, &post.Title, &post.Url, &post.Content, &post.Curriculum, &post.CurriculumId, &post.CreatedAt, &post.UpdatedAt); err != nil {
			return nil, fmt.Errorf("fail: rows.Scan, %v", err)
		}
	}

	likes, err := GetLikeCount(id)
	if err != nil {
		return nil, fmt.Errorf("fail: GetLikeCount, %v", err)
	}
	post.LikeCount = likes

	comments, err := GetCommentCount(id)
	if err != nil {
		return nil, fmt.Errorf("fail: GetCommentCount, %v", err)
	}
	post.CommentCount = comments

	bytes, err = json.Marshal(post)
	if err != nil {
		return nil, fmt.Errorf("fail: json.Marshal, %v", err)
	}
	return bytes, nil
}

func UpdatePost(p model.Post) (bytes []byte, err error) {
	tx, err := Db.Begin()
	if err != nil {
		return nil, fmt.Errorf("fail: Db.Begin, %v", err)
	}
	if _, err := tx.Exec("UPDATE posts "+
		"SET title = ?, url = ?, content = ?, curriculum_id = ? "+
		"WHERE id = ?",
		p.Title, p.Url, p.Content, p.CurriculumId, p.Id); err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			return nil, fmt.Errorf("fail: db.Rollback, %v", rollbackErr)
		}
		return nil, fmt.Errorf("fail: db.Exec, %v", err)
	}
	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("fail: db.Commit, %v", err)
	}
	response := map[string]string{"post_id": p.Id}
	bytes, err = json.Marshal(response)
	if err != nil {
		return nil, fmt.Errorf("fail: json.Marshal, %v", err)
	}
	return bytes, nil
}

func DeletePost(id string) (bytes []byte, err error) {
	tx, err := Db.Begin()
	if err != nil {
		return nil, fmt.Errorf("fail: Db.Begin, %v", err)
	}
	if _, err := tx.Exec("DELETE FROM posts "+
		"WHERE id = ?", id); err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			return nil, fmt.Errorf("fail: db.Rollback, %v", rollbackErr)
		}
		return nil, fmt.Errorf("fail: db.Exec, %v", err)
	}
	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("fail: db.Commit, %v", err)
	}
	response := map[string]string{"message": "success"}
	bytes, err = json.Marshal(response)
	if err != nil {
		return nil, fmt.Errorf("fail: json.Marshal, %v", err)
	}
	return bytes, nil
}

func GetAllCurriculums() (bytes []byte, err error) {
	rows, err := Db.Query("SELECT id, curriculum " +
		"FROM curriculums")
	if err != nil {
		return nil, fmt.Errorf("fail: db.Query, %v", err)
	}
	defer rows.Close()

	curriculums := make([]model.Curriculum, 0)
	for rows.Next() {
		var c model.Curriculum
		if err := rows.Scan(&c.Id, &c.Curriculum); err != nil {
			return nil, fmt.Errorf("fail: rows.Scan, %v", err)
		}
		curriculums = append(curriculums, c)
	}
	bytes, err = json.Marshal(curriculums)
	if err != nil {
		return nil, fmt.Errorf("fail: json.Marshal, %v", err)
	}
	return bytes, nil
}

func GetAllTags() (bytes []byte, err error) {
	rows, err := Db.Query("SELECT id, name " +
		"FROM tag_names")
	if err != nil {
		return nil, fmt.Errorf("fail: db.Query, %v", err)
	}
	defer rows.Close()

	tags := make([]model.Tag, 0)
	for rows.Next() {
		var t model.Tag
		if err := rows.Scan(&t.Id, &t.Tag); err != nil {
			return nil, fmt.Errorf("fail: rows.Scan, %v", err)
		}
		tags = append(tags, t)
	}
	bytes, err = json.Marshal(tags)
	if err != nil {
		return nil, fmt.Errorf("fail: json.Marshal, %v", err)
	}
	return bytes, nil
}

func GetAllPostsByUser(id string) (bytes []byte, err error) {
	rows, err := Db.Query(
		"SELECT p.id, ca.category, u.name, p.title, p.url, p.content, cu.curriculum, p.created_at, p.updated_at "+
			"FROM posts p "+
			"INNER JOIN categories ca ON p.category_id = ca.id "+
			"INNER JOIN users u ON p.user_id = u.id "+
			"INNER JOIN curriculums cu ON p.curriculum_id = cu.id "+
			"WHERE p.user_id = ? "+
			"ORDER BY p.created_at", id)

	if err != nil {
		return nil, fmt.Errorf("fail: db.Query, %v", err)
	}
	defer rows.Close()

	posts := make([]model.GetPost, 0)
	for rows.Next() {
		var p model.GetPost
		if err := rows.Scan(&p.Id, &p.Category, &p.User, &p.Title, &p.Url, &p.Content, &p.Curriculum, &p.CreatedAt, &p.UpdatedAt); err != nil {
			return nil, fmt.Errorf("fail: rows.Scan, %v", err)
		}
		likes, err := GetLikeCount(id)
		if err != nil {
			return nil, fmt.Errorf("fail: GetLikeCount, %v", err)
		}
		p.LikeCount = likes
		comments, err := GetCommentCount(id)
		if err != nil {
			return nil, fmt.Errorf("fail: GetCommentCount, %v", err)
		}
		p.CommentCount = comments
		posts = append(posts, p)
	}

	bytes, err = json.Marshal(posts)
	if err != nil {
		return nil, fmt.Errorf("fail: json.Marshal, %v", err)
	}
	return bytes, nil
}

func GetAllPostsByTag(id string) (bytes []byte, err error) {
	rows, err := Db.Query("SELECT p.id, ca.category, u.name, u.id, p.title, p.content, cu.curriculum, p.created_at, p.updated_at "+
		"FROM posts p "+
		"INNER JOIN categories ca ON p.category_id = ca.id "+
		"INNER JOIN users u ON p.user_id = u.id "+
		"INNER JOIN curriculums cu ON p.curriculum_id = cu.id "+
		"INNER JOIN post_tags pt ON p.id = pt.post_id "+
		"INNER JOIN tag_names tn ON pt.tag_id = tn.id "+
		"WHERE tn.id = ? "+
		"ORDER BY p.created_at DESC ", id)
	if err != nil {
		return nil, fmt.Errorf("fail: db.Query, %v", err)
	}
	defer rows.Close()

	posts := make([]model.GetPost, 0)
	for rows.Next() {
		var p model.GetPost
		if err := rows.Scan(&p.Id, &p.Category, &p.User, &p.UserId, &p.Title, &p.Url, &p.Content, &p.Curriculum, &p.CreatedAt, &p.UpdatedAt); err != nil {
			return nil, fmt.Errorf("fail: rows.Scan, %v", err)
		}
		likes, err := GetLikeCount(p.Id)
		if err != nil {
			return nil, fmt.Errorf("fail: GetLikeCount, %v", err)
		}
		p.LikeCount = likes
		comments, err := GetCommentCount(p.Id)
		if err != nil {
			return nil, fmt.Errorf("fail: GetCommentCount, %v", err)
		}
		p.CommentCount = comments
		posts = append(posts, p)
	}
	bytes, err = json.Marshal(posts)
	if err != nil {
		return nil, fmt.Errorf("fail: json.Marshal, %v", err)
	}
	return bytes, nil
}

func GetAllPostsByCategory(id int) (bytes []byte, err error) {
	rows, err := Db.Query("SELECT p.id, ca.category, u.name, u.id, p.title, p.url, p.content, cu.curriculum, p.created_at, p.updated_at "+
		"FROM posts p "+
		"INNER JOIN categories ca ON p.category_id = ca.id "+
		"INNER JOIN users u ON p.user_id = u.id "+
		"INNER JOIN curriculums cu ON p.curriculum_id = cu.id "+
		"WHERE p.category_id = ? "+
		"ORDER BY p.created_at DESC", id)

	if err != nil {
		return nil, fmt.Errorf("fail: db.Query, %v", err)
	}
	defer rows.Close()

	posts := make([]model.GetPost, 0)
	for rows.Next() {
		var p model.GetPost
		if err := rows.Scan(&p.Id, &p.Category, &p.User, &p.UserId, &p.Title, &p.Url, &p.Content, &p.Curriculum, &p.CreatedAt, &p.UpdatedAt); err != nil {
			return nil, fmt.Errorf("fail: rows.Scan, %v", err)
		}
		likes, err := GetLikeCount(p.Id)
		if err != nil {
			return nil, fmt.Errorf("fail: GetLikeCount, %v", err)
		}
		p.LikeCount = likes
		comments, err := GetCommentCount(p.Id)
		if err != nil {
			return nil, fmt.Errorf("fail: GetCommentCount, %v", err)
		}
		p.CommentCount = comments
		posts = append(posts, p)
	}
	bytes, err = json.Marshal(posts)
	if err != nil {
		return nil, fmt.Errorf("fail: json.Marshal, %v", err)
	}
	return bytes, nil
}

func GetAllPostsByCurriculum(id string) (bytes []byte, err error) {
	rows, err := Db.Query("SELECT p.id, ca.category, u.name, u.id, p.title, p.url, p.content, cu.curriculum, p.created_at, p.updated_at "+
		"FROM posts p "+
		"INNER JOIN categories ca ON p.category_id = ca.id "+
		"INNER JOIN users u ON p.user_id = u.id "+
		"INNER JOIN curriculums cu ON p.curriculum_id = cu.id "+
		"WHERE p.curriculum_id = ? "+
		"ORDER BY p.created_at DESC", id)
	if err != nil {
		return nil, fmt.Errorf("fail: db.Query, %v", err)
	}
	defer rows.Close()

	posts := make([]model.GetPost, 0)
	for rows.Next() {
		var p model.GetPost
		if err := rows.Scan(&p.Id, &p.Category, &p.User, &p.UserId, &p.Title, &p.Url, &p.Content, &p.Curriculum, &p.CreatedAt, &p.UpdatedAt); err != nil {
			return nil, fmt.Errorf("fail: rows.Scan, %v", err)
		}
		likes, err := GetLikeCount(p.Id)
		if err != nil {
			return nil, fmt.Errorf("fail: GetLikeCount, %v", err)
		}
		p.LikeCount = likes
		comments, err := GetCommentCount(p.Id)
		if err != nil {
			return nil, fmt.Errorf("fail: GetCommentCount, %v", err)
		}
		p.CommentCount = comments
		posts = append(posts, p)
	}
	bytes, err = json.Marshal(posts)
	if err != nil {
		return nil, fmt.Errorf("fail: json.Marshal, %v", err)
	}
	return bytes, nil
}

func GetAllPostsByDate() (bytes []byte, err error) {
	rows, err := Db.Query("SELECT p.id, ca.category, u.name, u.id, p.title, p.url, p.content, cu.curriculum, p.created_at, p.updated_at " +
		"FROM posts p " +
		"INNER JOIN categories ca ON p.category_id = ca.id " +
		"INNER JOIN users u ON p.user_id = u.id " +
		"INNER JOIN curriculums cu ON p.curriculum_id = cu.id " +
		"ORDER BY p.created_at DESC")

	if err != nil {
		return nil, fmt.Errorf("fail: db.Query, %v", err)
	}
	defer rows.Close()

	posts := make([]model.GetPost, 0)
	for rows.Next() {
		var p model.GetPost
		if err := rows.Scan(&p.Id, &p.Category, &p.User, &p.UserId, &p.Title, &p.Url, &p.Content, &p.Curriculum, &p.CreatedAt, &p.UpdatedAt); err != nil {
			return nil, fmt.Errorf("fail: rows.Scan, %v", err)
		}
		likes, err := GetLikeCount(p.Id)
		if err != nil {
			return nil, fmt.Errorf("fail: GetLikeCount, %v", err)
		}
		p.LikeCount = likes
		comments, err := GetCommentCount(p.Id)
		if err != nil {
			return nil, fmt.Errorf("fail: GetCommentCount, %v", err)
		}
		p.CommentCount = comments
		posts = append(posts, p)

	}
	bytes, err = json.Marshal(posts)
	if err != nil {
		return nil, fmt.Errorf("fail: json.Marshal, %v", err)
	}
	return bytes, nil
}

func LikePost(l model.Like) (bytes []byte, err error) {
	tx, err := Db.Begin()
	if err != nil {
		return nil, fmt.Errorf("fail: Db.Begin, %v", err)
	}
	if _, err := tx.Exec("INSERT INTO likes (id, post_id, user_id) "+
		"VALUES (?, ?, ?)",
		l.Id, l.PostId, l.UserId); err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			return nil, fmt.Errorf("fail: db.Rollback, %v", rollbackErr)
		}
		return nil, fmt.Errorf("fail: db.Exec, %v", err)
	}
	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("fail: db.Commit, %v", err)
	}
	response := map[string]string{"message": "success"}
	bytes, err = json.Marshal(response)
	if err != nil {
		return nil, fmt.Errorf("fail: json.Marshal, %v", err)
	}
	return bytes, nil
}

func UnlikePost(l model.Unlike) (bytes []byte, err error) {
	tx, err := Db.Begin()
	if err != nil {
		return nil, fmt.Errorf("fail: Db.Begin, %v", err)
	}
	if _, err := tx.Exec("DELETE FROM likes "+
		"WHERE user_id = ? AND post_id = ?", l.UserId, l.PostId); err != nil {
		err := tx.Rollback()
		if err != nil {
			return nil, fmt.Errorf("fail: db.Rollback, %v", err)
		}
		return nil, fmt.Errorf("fail: db.Exec, %v", err)
	}
	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("fail: db.Commit, %v", err)
	}
	response := map[string]string{"message": "success"}
	bytes, err = json.Marshal(response)
	if err != nil {
		return nil, fmt.Errorf("fail: json.Marshal, %v", err)
	}
	return bytes, nil
}

func GetLikeCount(id string) (count int, err error) {
	rows, err := Db.Query("SELECT COUNT(*) "+
		"FROM likes "+
		"WHERE post_id = ?", id)
	if err != nil {
		return 0, fmt.Errorf("fail: db.Query, %v", err)
	}
	for rows.Next() {
		if err := rows.Scan(&count); err != nil {
			if err := rows.Close(); err != nil {
				return 0, fmt.Errorf("fail: rows.Close(), %v", err)
			}
			return 0, fmt.Errorf("fail: rows.Scan, %v", err)
		}
	}
	return count, nil
}

func GetLikedPosts(id string) (bytes []byte, err error) {
	rows, err := Db.Query("SELECT p.id, ca.category, u.name, p.title, p.url, p.content, cu.curriculum, p.created_at, p.updated_at "+
		"FROM posts p "+
		"INNER JOIN categories ca ON p.category_id = ca.id "+
		"INNER JOIN users u ON p.user_id = u.id "+
		"INNER JOIN curriculums cu ON p.curriculum_id = cu.id "+
		"INNER JOIN likes l ON p.id = l.post_id "+
		"WHERE l.user_id = ? "+
		"ORDER BY l.liked_at DESC", id)

	if err != nil {
		return nil, fmt.Errorf("fail: db.Query, %v", err)
	}
	posts := make([]model.GetPost, 0)
	for rows.Next() {
		var p model.GetPost
		if err := rows.Scan(&p.Id, &p.Category, &p.User, &p.Title, &p.Url, &p.Content, &p.Curriculum, &p.CreatedAt, &p.UpdatedAt); err != nil {
			if err := rows.Close(); err != nil {
				return nil, fmt.Errorf("fail: rows.Close(), %v", err)
			}
			return nil, fmt.Errorf("fail: rows.Scan, %v", err)
		}
		likes, err := GetLikeCount(id)
		if err != nil {
			return nil, fmt.Errorf("fail: GetLikeCount, %v", err)
		}
		p.LikeCount = likes
		comments, err := GetCommentCount(id)
		if err != nil {
			return nil, fmt.Errorf("fail: GetCommentCount, %v", err)
		}
		p.CommentCount = comments
		posts = append(posts, p)
	}
	bytes, err = json.Marshal(posts)
	if err != nil {
		return nil, fmt.Errorf("fail: json.Marshal, %v", err)
	}
	return bytes, nil
}

func CommentPost(c model.Comment) (bytes []byte, err error) {
	tx, err := Db.Begin()
	if err != nil {
		return nil, fmt.Errorf("fail: Db.Begin, %v", err)
	}
	if _, err := tx.Exec("INSERT INTO comments (id, post_id, user_id, content) "+
		"VALUES (?, ?, ?, ?)",
		c.Id, c.PostId, c.UserId, c.Content); err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			return nil, fmt.Errorf("fail: db.Rollback, %v", rollbackErr)
		}
		return nil, fmt.Errorf("fail: db.Exec, %v", err)
	}
	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("fail: db.Commit, %v", err)
	}
	bytes, err = json.Marshal("success")
	if err != nil {
		return nil, fmt.Errorf("fail: json.Marshal, %v", err)
	}
	return bytes, nil
}

func DeleteComment(id string) (bytes []byte, err error) {
	tx, err := Db.Begin()
	if err != nil {
		return nil, fmt.Errorf("fail: Db.Begin, %v\n", err)
	}
	if _, err := tx.Exec("DELETE FROM comments "+
		"WHERE id = ?", id); err != nil {
		err := tx.Rollback()
		if err != nil {
			return nil, fmt.Errorf("fail: db.Rollback, %v\n", err)
		}
		return nil, fmt.Errorf("fail: db.Exec, %v\n", err)
	}
	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("fail: db.Commit, %v\n", err)
	}
	response := map[string]string{"message": "success"}
	bytes, err = json.Marshal(response)
	if err != nil {
		return nil, fmt.Errorf("fail: json.Marshal, %v", err)
	}
	return bytes, nil
}

func UpdateComment(c model.Comment) (bytes []byte, err error) {
	tx, err := Db.Begin()
	if err != nil {
		return nil, fmt.Errorf("fail: Db.Begin, %v\n", err)
	}
	if _, err := tx.Exec("UPDATE comments "+
		"SET content = ? "+
		"WHERE id = ?",
		c.Content, c.Id); err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			return nil, fmt.Errorf("fail: db.Rollback, %v", rollbackErr)
		}
		return nil, fmt.Errorf("fail: db.Exec, %v\n", err)
	}
	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("fail: db.Commit, %v\n", err)
	}
	response := map[string]string{"message": "success"}
	bytes, err = json.Marshal(response)
	if err != nil {
		return nil, fmt.Errorf("fail: json.Marshal, %v", err)
	}
	return bytes, nil
}

func GetAllCommentsByPost(id string) (bytes []byte, err error) {
	rows, err := Db.Query("SELECT c.id, u.name, u.id, u.image, c.content, c.created_at "+
		"FROM comments c "+
		"INNER JOIN users u ON c.user_id = u.id "+
		"WHERE c.post_id = ? "+
		"ORDER BY c.created_at;", id)
	if err != nil {
		return nil, fmt.Errorf("fail: db.Query, %v\n", err)
	}
	comments := make([]model.GetComment, 0)
	for rows.Next() {
		var c model.GetComment
		if err := rows.Scan(&c.Id, &c.User, &c.UserId, &c.UserImage, &c.Content, &c.CreatedAt); err != nil {
			if err := rows.Close(); err != nil {
				return nil, fmt.Errorf("fail: rows.Close(), %v\n", err)
			}
			return nil, fmt.Errorf("fail: rows.Scan, %v\n", err)
		}
		comments = append(comments, c)
	}
	bytes, err = json.Marshal(comments)
	if err != nil {
		return nil, fmt.Errorf("fail: json.Marshal, %v\n", err)
	}
	return bytes, nil
}

func GetCommentCount(id string) (count int, err error) {
	rows, err := Db.Query("SELECT COUNT(*) "+
		"FROM comments "+
		"WHERE post_id = ?", id)
	if err != nil {
		return 0, fmt.Errorf("fail: db.Query, %v", err)
	}
	for rows.Next() {
		if err := rows.Scan(&count); err != nil {
			if err := rows.Close(); err != nil {
				return 0, fmt.Errorf("fail: rows.Close(), %v", err)
			}
			return 0, fmt.Errorf("fail: rows.Scan, %v", err)
		}
	}
	return count, nil
}
