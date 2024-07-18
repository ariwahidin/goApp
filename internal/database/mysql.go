package database

import (
	"database/sql"
	"fmt"
	"goapp/config"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@/%s",
		config.DBUser, config.DBPassword, config.DBName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
