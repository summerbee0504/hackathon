package dao

import (
	"database/sql"
	"fmt"
	"log"
)

var Db *sql.DB

func Init() {

	//if err := godotenv.Load(".env_mysql"); err != nil {
	//	log.Fatalf("fail: godotenv.Load, %v\n", err)
	//}
	//mysqlUser := os.Getenv("MYSQL_USER")
	//mysqlPwd := os.Getenv("MYSQL_PWD")
	//mysqlHost := os.Getenv("MYSQL_HOST")
	//mysqlDatabase := os.Getenv("MYSQL_DATABASE")

	mysqlUser := "test_user"
	mysqlPwd := "password"
	mysqlDatabase := "hackathon_local"

	//connStr := fmt.Sprintf("%s:%s@%s/%s", mysqlUser, mysqlPwd, mysqlHost, mysqlDatabase)
	//_db, err := sql.Open("mysql", connStr)
	_db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@(localhost:3306)/%s", mysqlUser, mysqlPwd, mysqlDatabase))
	if err != nil {
		log.Fatalf("fail: sql.Open, %v\n", err)
	}

	if err := _db.Ping(); err != nil {
		log.Fatalf("fail: _db.Ping, %v\n", err)
	}
	Db = _db
}

func CloseDb() (err error) {
	err = Db.Close()
	return err
}
