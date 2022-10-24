package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func Connection() *sql.DB {
	db, err := sql.Open("mysql", "root:root@/mydb?charset=utf8&parseTime=True")
	if err != nil {
		err := fmt.Errorf("open connection: %w", err)
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		err := fmt.Errorf("ping connection test: %w", err)
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db
}
