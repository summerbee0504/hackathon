package dao

import (
	"encoding/json"
	"fmt"
	"hackathon/model"
)

func MakePost(p model.Post) error {
	tx, err := Db.Begin()
	if err != nil {
		return fmt.Errorf("fail: Db.Begin, %v", err)
	}
	if _, err := tx.Exec("INSERT INTO posts (id, category_id, user_id, title, content, curriculum_id) "+
		"VALUES (?, ?, ?, ?, ?, ?)",
		p.Id, p.CategoryId, p.UserId, p.Title, p.Content, p.CurriculumId); err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			return fmt.Errorf("fail: db.Rollback, %v", rollbackErr)
		}
		return fmt.Errorf("fail: db.Exec, %v", err)
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("fail: db.Commit, %v", err)
	}
	return nil
}

func GetPost(id string) (bytes []byte, err error) {
	rows, err := Db.Query("SELECT ca.category, u.name, p.title, p.content, cu.curriculum, p.created_at, p.updated_at "+
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
		if err := rows.Scan(&post.Category, &post.User, &post.Title, &post.Content, &post.Curriculum, &post.CreatedAt, &post.UpdatedAt); err != nil {
			return nil, fmt.Errorf("fail: rows.Scan, %v", err)
		}
	}
	bytes, err = json.Marshal(post)
	if err != nil {
		return nil, fmt.Errorf("fail: json.Marshal, %v", err)
	}
	return bytes, nil
}

func UpdatePost(p model.Post) error {
	tx, err := Db.Begin()
	if err != nil {
		return fmt.Errorf("fail: Db.Begin, %v", err)
	}
	if _, err := tx.Exec("UPDATE posts "+
		"SET category_id = ?, user_id = ?, title = ?, content = ?, curriculum_id = ? "+
		"WHERE id = ?",
		p.CategoryId, p.UserId, p.Title, p.Content, p.CurriculumId, p.Id); err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			return fmt.Errorf("fail: db.Rollback, %v", rollbackErr)
		}
		return fmt.Errorf("fail: db.Exec, %v", err)
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("fail: db.Commit, %v", err)
	}
	return nil
}

func DeletePost(id string) error {
	tx, err := Db.Begin()
	if err != nil {
		return fmt.Errorf("fail: Db.Begin, %v", err)
	}
	if _, err := tx.Exec("DELETE FROM posts "+
		"WHERE id = ?", id); err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			return fmt.Errorf("fail: db.Rollback, %v", rollbackErr)
		}
		return fmt.Errorf("fail: db.Exec, %v", err)
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("fail: db.Commit, %v", err)
	}
	return nil
}

func GetAllPostsByUser(id string) (bytes []byte, err error) {
	rows, err := Db.Query("SELECT p.id, ca.category, u.name, p.title, p.updated_at "+
		"FROM posts p "+
		"INNER JOIN categories ca ON p.category_id = ca.id "+
		"INNER JOIN users u ON p.user_id = u.id "+
		"WHERE p.user_id = ?", id)
	if err != nil {
		return nil, fmt.Errorf("fail: db.Query, %v", err)
	}
	defer rows.Close()

	posts := make([]model.GetPost, 0)
	for rows.Next() {
		var p model.GetPost
		if err := rows.Scan(&p.Id, &p.Category, &p.User, &p.Title, &p.UpdatedAt); err != nil {
			return nil, fmt.Errorf("fail: rows.Scan, %v", err)
		}
		posts = append(posts, p)
	}
	bytes, err = json.Marshal(posts)
	if err != nil {
		return nil, fmt.Errorf("fail: json.Marshal, %v", err)
	}
	return bytes, nil
}

func GetAllPostsByTag(id string) (bytes []byte, err error) {
	rows, err := Db.Query("SELECT p.id, t.name, ca.category, u.name, p.title, p.updated_at "+
		"FROM posts p "+
		"INNER JOIN categories ca ON p.category_id = ca.id "+
		"INNER JOIN users u ON p.user_id = u.id "+
		"INNER JOIN tags t ON p.tag_id = t.id "+
		"WHERE p.tag_id = ?", id)
	if err != nil {
		return nil, fmt.Errorf("fail: db.Query, %v", err)
	}
	defer rows.Close()

	posts := make([]model.GetPost, 0)
	for rows.Next() {
		var p model.GetPost
		if err := rows.Scan(&p.Id, &p.Category, &p.User, &p.Title, &p.UpdatedAt); err != nil {
			return nil, fmt.Errorf("fail: rows.Scan, %v", err)
		}
		posts = append(posts, p)
	}
	bytes, err = json.Marshal(posts)
	if err != nil {
		return nil, fmt.Errorf("fail: json.Marshal, %v", err)
	}
	return bytes, nil
}

func GetAllPostsByCategory(id int) (bytes []byte, err error) {
	rows, err := Db.Query("SELECT p.id, ca.category, u.name, p.title, p.updated_at "+
		"FROM posts p "+
		"INNER JOIN categories ca ON p.category_id = ca.id "+
		"INNER JOIN users u ON p.user_id = u.id "+
		"WHERE p.category_id = ? "+
		"ORDER BY p.updated_at", id)
	if err != nil {
		return nil, fmt.Errorf("fail: db.Query, %v", err)
	}
	defer rows.Close()

	posts := make([]model.GetPost, 0)
	for rows.Next() {
		var p model.GetPost
		if err := rows.Scan(&p.Id, &p.Category, &p.User, &p.Title, &p.UpdatedAt); err != nil {
			return nil, fmt.Errorf("fail: rows.Scan, %v", err)
		}
		posts = append(posts, p)
	}
	bytes, err = json.Marshal(posts)
	if err != nil {
		return nil, fmt.Errorf("fail: json.Marshal, %v", err)
	}
	return bytes, nil
}

func GetAllPostsByCurriculum(id int) (bytes []byte, err error) {
	rows, err := Db.Query("SELECT p.id, ca.category, u.name, p.title, p.updated_at "+
		"FROM posts p "+
		"INNER JOIN categories ca ON p.category_id = ca.id "+
		"INNER JOIN users u ON p.user_id = u.id "+
		"INNER JOIN curriculums cu ON p.curriculum_id = cu.id "+
		"WHERE p.curriculum_id = ? "+
		"ORDER BY p.updated_at", id)
	if err != nil {
		return nil, fmt.Errorf("fail: db.Query, %v", err)
	}
	defer rows.Close()

	posts := make([]model.GetPost, 0)
	for rows.Next() {
		var p model.GetPost
		if err := rows.Scan(&p.Id, &p.Category, &p.User, &p.Title, &p.UpdatedAt); err != nil {
			return nil, fmt.Errorf("fail: rows.Scan, %v", err)
		}
		posts = append(posts, p)
	}
	bytes, err = json.Marshal(posts)
	if err != nil {
		return nil, fmt.Errorf("fail: json.Marshal, %v", err)
	}
	return bytes, nil
}

func GetAllPostsByDate() (bytes []byte, err error) {
	rows, err := Db.Query("SELECT p.id, ca.category, u.name, p.title, p.updated_at " +
		"FROM posts p " +
		"INNER JOIN categories ca ON p.category_id = ca.id " +
		"INNER JOIN users u ON p.user_id = u.id " +
		"ORDER BY p.updated_at")
	if err != nil {
		return nil, fmt.Errorf("fail: db.Query, %v", err)
	}
	defer rows.Close()

	posts := make([]model.GetPost, 0)
	for rows.Next() {
		var p model.GetPost
		if err := rows.Scan(&p.Id, &p.Category, &p.User, &p.Title, &p.UpdatedAt); err != nil {
			return nil, fmt.Errorf("fail: rows.Scan, %v", err)
		}
		posts = append(posts, p)
	}
	bytes, err = json.Marshal(posts)
	if err != nil {
		return nil, fmt.Errorf("fail: json.Marshal, %v", err)
	}
	return bytes, nil
}

func LikePost(l model.Like) error {
	tx, err := Db.Begin()
	if err != nil {
		return fmt.Errorf("fail: Db.Begin, %v", err)
	}
	if _, err := tx.Exec("INSERT INTO likes (id, post_id, user_id) "+
		"VALUES (?, ?, ?)",
		l.Id, l.PostId, l.UserId); err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			return fmt.Errorf("fail: db.Rollback, %v", rollbackErr)
		}
		return fmt.Errorf("fail: db.Exec, %v", err)
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("fail: db.Commit, %v", err)
	}
	return nil
}

func UnlikePost(id string) error {
	tx, err := Db.Begin()
	if err != nil {
		return fmt.Errorf("fail: Db.Begin, %v", err)
	}
	if _, err := tx.Exec("DELETE FROM likes "+
		"WHERE id = ?", id); err != nil {
		err := tx.Rollback()
		if err != nil {
			return fmt.Errorf("fail: db.Rollback, %v", err)
		}
		return fmt.Errorf("fail: db.Exec, %v", err)
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("fail: db.Commit, %v", err)
	}
	return nil
}

func GetLikeCount(id string) (bytes []byte, err error) {
	rows, err := Db.Query("SELECT COUNT(*) "+
		"FROM likes "+
		"WHERE post_id = ?", id)
	if err != nil {
		return nil, fmt.Errorf("fail: db.Query, %v", err)
	}
	var count int
	for rows.Next() {
		if err := rows.Scan(&count); err != nil {
			if err := rows.Close(); err != nil {
				return nil, fmt.Errorf("fail: rows.Close(), %v", err)
			}
			return nil, fmt.Errorf("fail: rows.Scan, %v", err)
		}
	}
	bytes, err = json.Marshal(count)
	if err != nil {
		return nil, fmt.Errorf("fail: json.Marshal, %v", err)
	}
	return bytes, nil
}

func GetLikedPosts(id string) (bytes []byte, err error) {
	rows, err := Db.Query("SELECT p.id, p.title, u.name, p.updated_at "+
		"FROM posts p "+
		"INNER JOIN likes l ON p.id = l.post_id "+
		"INNER JOIN users u ON p.user_id = u.id "+
		"WHERE l.user_id = ? "+
		"ORDER BY l.liked_at", id)
	if err != nil {
		return nil, fmt.Errorf("fail: db.Query, %v", err)
	}
	posts := make([]model.GetPost, 0)
	for rows.Next() {
		var p model.GetPost
		if err := rows.Scan(&p.Id, &p.Title, &p.User, &p.UpdatedAt); err != nil {
			if err := rows.Close(); err != nil {
				return nil, fmt.Errorf("fail: rows.Close(), %v", err)
			}
			return nil, fmt.Errorf("fail: rows.Scan, %v", err)
		}
		posts = append(posts, p)
	}
	bytes, err = json.Marshal(posts)
	if err != nil {
		return nil, fmt.Errorf("fail: json.Marshal, %v", err)
	}
	return bytes, nil
}

func CommentPost(c model.Comment) error {
	tx, err := Db.Begin()
	if err != nil {
		return fmt.Errorf("fail: Db.Begin, %v", err)
	}
	if _, err := tx.Exec("INSERT INTO comments (id, post_id, user_id, content)"+
		"VALUES (?, ?, ?, ?)",
		c.Id, c.PostId, c.UserId, c.Content); err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			return fmt.Errorf("fail: db.Rollback, %v", rollbackErr)
		}
		return fmt.Errorf("fail: db.Exec, %v", err)
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("fail: db.Commit, %v", err)
	}
	return nil
}

func DeleteComment(id string) error {
	tx, err := Db.Begin()
	if err != nil {
		return fmt.Errorf("fail: Db.Begin, %v\n", err)
	}
	if _, err := tx.Exec("DELETE FROM comments "+
		"WHERE id = ?", id); err != nil {
		err := tx.Rollback()
		if err != nil {
			return fmt.Errorf("fail: db.Rollback, %v\n", err)
		}
		return fmt.Errorf("fail: db.Exec, %v\n", err)
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("fail: db.Commit, %v\n", err)
	}
	return nil
}

func UpdateComment(c model.Comment) error {
	tx, err := Db.Begin()
	if err != nil {
		return fmt.Errorf("fail: Db.Begin, %v\n", err)
	}
	if _, err := tx.Exec("UPDATE comments "+
		"SET content = ? "+
		"WHERE id = ?",
		c.Content, c.Id); err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			return fmt.Errorf("fail: db.Rollback, %v", rollbackErr)
		}
		return fmt.Errorf("fail: db.Exec, %v\n", err)
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("fail: db.Commit, %v\n", err)
	}
	return nil
}

func GetAllCommentsByPost(id string) (bytes []byte, err error) {
	rows, err := Db.Query("SELECT c.id, u.name, c.content, c.updated_at "+
		"FROM comments c "+
		"INNER JOIN users u ON c.user_id = u.id "+
		"WHERE c.post_id = ? "+
		"ORDER BY c.updated_at", id)
	if err != nil {
		return nil, fmt.Errorf("fail: db.Query, %v\n", err)
	}
	comments := make([]model.GetComment, 0)
	for rows.Next() {
		var c model.GetComment
		if err := rows.Scan(&c.Id, &c.User, &c.Content, &c.UpdatedAt); err != nil {
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
