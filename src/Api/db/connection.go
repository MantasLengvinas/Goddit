package db

import (
	"log"

	"github.com/rqlite/gorqlite"
)

func InitDB() (*gorqlite.Connection, error) {
	conn, err := gorqlite.Open("http://")

	if err != nil {
		log.Fatalln("DB connection cannot be established")
		log.Fatalln(err.Error())
		return nil, err
	}

	return conn, nil
}

func CloseDB(conn *gorqlite.Connection) {
	conn.Close()
}
