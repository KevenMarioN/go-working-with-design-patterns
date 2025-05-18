package main

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	maxOpenDBConn = 25
	maxIDLEDBConn = 25
	maxDBLifeTime = 5 * time.Minute
)

func initMySQLB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenDBConn)
	db.SetMaxIdleConns(maxIDLEDBConn)
	db.SetConnMaxLifetime(maxDBLifeTime)

	return db, nil
}
