package dao

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"hackathon/model"
	"log"
	"net/http"
	"os"
)

var db *sql.DB

func Init() {

	if err := godotenv.Load(".env_mysql"); err != nil {
		log.Fatalf("fail: godotenv.Load, %v\n", err)
	}
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPwd := os.Getenv("MYSQL_PWD")
	mysqlHost := os.Getenv("MYSQL_HOST")
	mysqlDatabase := os.Getenv("MYSQL_DATABASE")

	connStr := fmt.Sprintf("%s:%s@%s/%s", mysqlUser, mysqlPwd, mysqlHost, mysqlDatabase)
	_db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatalf("fail: sql.Open, %v\n", err)
	}

	if err := _db.Ping(); err != nil {
		log.Fatalf("fail: _db.Ping, %v\n", err)
	}
	db = _db
}

func SearchUser(w http.ResponseWriter, name string) (bytes []byte) {
	rows, err := db.Query("SELECT id, name, age FROM user WHERE name = ?", name)
	if err != nil {
		log.Printf("fail: db.Query, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	users := make([]model.User, 0)
	for rows.Next() {
		var u model.User
		if err := rows.Scan(&u.Id, &u.Name, &u.Age); err != nil {
			log.Printf("fail: rows.Scan, %v\n", err)

			if err := rows.Close(); err != nil {
				log.Printf("fail: rows.Close(), %v\n", err)
			}
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		users = append(users, u)
	}

	bytes, err = json.Marshal(users)
	if err != nil {
		log.Printf("fail: json.Marshal, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	return bytes
}

func RegisterUser(w http.ResponseWriter, u model.User, id string) {
	tx, err := db.Begin()
	if err != nil {
		log.Printf("fail: db.Begin, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if _, err := tx.Exec("INSERT INTO user (id, name, age) VALUES (?, ?, ?)", id, u.Name, u.Age); err != nil {
		err := tx.Rollback()
		if err != nil {
			log.Printf("fail: db.Rollback, %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		log.Printf("fail: db.Exec, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := tx.Commit(); err != nil {
		log.Printf("fail: db.Commit, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func CloseDb() (err error) {
	err = db.Close()
	return err
}
