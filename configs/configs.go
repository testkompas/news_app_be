package configs

import (
	"fmt"
	"os"
)

const (
	HASHKEY = "c5dda7a77f7dc8e29cd2d949ccc201c02e1afdd5d4a44993d2a81509d53c6954"

	MYSQL_USER     = "user123"
	MYSQL_PASSWORD = "2022Mysql!!"
	MYSQL_HOST     = "172.17.0.1"
	MYSQL_PORT     = "3306"

	DB_NAME = "news_app"
)

func GetMySqlDSN() (result string) {

	author := os.Getenv("MYSQL_USER")
	if author == "" {
		author = MYSQL_USER
	}
	password := os.Getenv("MYSQL_PASS")
	if password == "" {
		password = MYSQL_PASSWORD
	}
	host := os.Getenv("MYSQL_HOST")
	if host == "" {
		host = MYSQL_HOST
	}
	port := os.Getenv("MYSQL_PORT")
	if port == "" {
		port = MYSQL_PORT
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = DB_NAME
	}

	result = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", author, password, host, port, dbName)
	return
}
