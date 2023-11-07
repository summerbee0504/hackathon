package dao

import (
	"encoding/json"
	"fmt"
	"hackathon/model"
)

// ShowUserList ユーザーリストを取得する関数
func ShowUserList() (bytes []byte, err error) {
	rows, err := Db.Query("SELECT u.id, u.name, u.term, p.permission " +
		"FROM users u " +
		"INNER JOIN permissions p ON u.permission_id = p.id " +
		"ORDER BY u.permission_id;")
	if err != nil {
		return nil, fmt.Errorf("fail: Db.Query, %v\n", err)
	}
	defer rows.Close()

	users := make([]model.ShowUserList, 0)
	for rows.Next() {
		var u model.ShowUserList
		if err := rows.Scan(&u.Id, &u.Name, &u.Term, &u.Permission); err != nil {
			return nil, fmt.Errorf("fail: rows.Scan, %v\n", err)
		}
		users = append(users, u)
	}

	bytes, err = json.Marshal(users)
	if err != nil {
		return nil, fmt.Errorf("fail: json.Marshal, %v\n", err)
	}
	return bytes, nil
}

// ShowUserDetail ユーザー詳細を取得する関数
func ShowUserDetail(id string) (bytes []byte, err error) {
	rows, err := Db.Query("SELECT u.id, u.name, u.term, u.bio, p.permission "+
		"FROM users u "+
		"INNER JOIN permissions p ON u.permission_id = p.id "+
		"WHERE u.id = ?;", id)
	if err != nil {
		return nil, fmt.Errorf("fail: Db.Query, %v\n", err)
	}
	defer rows.Close()

	var u model.ShowUser

	for rows.Next() {
		if err := rows.Scan(&u.Id, &u.Name, &u.Term, &u.Bio, &u.Permission); err != nil {
			return nil, fmt.Errorf("fail: rows.Scan, %v\n", err)
		}
	}

	bytes, err = json.Marshal(u)
	if err != nil {
		return nil, fmt.Errorf("fail: json.Marshal, %v", err)
	}

	return bytes, nil
}

// RegisterUser ユーザーを登録する関数
func RegisterUser(u model.User) (bytes []byte, err error) {
	tx, err := Db.Begin()
	if err != nil {
		return nil, fmt.Errorf("fail: db.Begin, %v", err)
	}

	_, err = tx.Exec("INSERT INTO users (id, name, term, bio, image, permission_id) VALUES (?, ?, ?, ?, ?, ?)",
		u.Id, u.Name, u.Term, u.Bio, u.Image, u.PermissionId)
	if err != nil {
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

// UpdateUser ユーザー情報を更新する関数
func UpdateUser(u model.UpdateUserDetails) (bytes []byte, err error) {
	tx, err := Db.Begin()
	if err != nil {
		return nil, fmt.Errorf("fail: db.Begin, %v", err)
	}

	_, err = tx.Exec("UPDATE users "+
		"SET name = ?, term = ?, bio = ? "+
		"WHERE id = ?", u.Name, u.Term, u.Bio, u.Id)
	if err != nil {
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

// DeleteUser ユーザーを削除する関数
func DeleteUser(id string) (bytes []byte, err error) {
	tx, err := Db.Begin()
	if err != nil {
		return nil, fmt.Errorf("fail: db.Begin, %v", err)
	}

	_, err = tx.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
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
