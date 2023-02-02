package app

import (
	"database/sql"
	"time"

	"github.com/Tiorzfk/d5c08d914cbd25fe7001490bbc7d24ed/helper"
	_ "github.com/go-sql-driver/mysql"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:bismilah123@tcp(localhost:3306)/shopping_cart")
	helper.PanicIfError(err)

	db.SetConnMaxIdleTime(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
