package config

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)


func Hello() *sql.DB {
	db, err := sql.Open("postgres", `host=localhost port=5432 user=postgres password=postgres dbname=DB_KA_MAKASSAR sslmode=disable`)
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)
	// defer db.Close()
	// fmt.Printf("SUCCESS CONNECT")
	return db
}