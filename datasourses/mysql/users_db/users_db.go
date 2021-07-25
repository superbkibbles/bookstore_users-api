package users_db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/superbkibbles/bookstore_utils-go/logger"
)

const (
	MYSQL_USER_USERNAME = "MYSQL_USER_USERNAME"
	MYSQL_USER_PASSWORD = "MYSQL_USER_PASSWORD"
	MYSQL_USER_HOST     = "MYSQL_USER_HOST"
	MYSQL_USER_SCHEMA   = "MYSQL_USER_SCHEMA"
)

var (
	Client *sql.DB

	username = os.Getenv(MYSQL_USER_USERNAME)
	password = os.Getenv(MYSQL_USER_PASSWORD)
	host     = os.Getenv(MYSQL_USER_HOST)
	schema   = os.Getenv(MYSQL_USER_SCHEMA)
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username, password, host, schema,
	)
	// dataSourceName := "root:A3201888118a@/users_db?charset=utf8"
	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}

	mysql.SetLogger(logger.Getlogger())
}
