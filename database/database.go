package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var db sql.DB*

func Pool() error {
	var err error
	if (db == nil){
		db, err := sql.Open("postgres", "user=devadmin password=devadmin123 port=5432 host=dev-db.cn2axrsupztc.ap-southeast-1.rds.amazonaws.com dbname=dev-db  sslmode=disable")
		if err != nil {
			log.Fatal("Error: Could not establish a connection with the database")
		}

		err = db.Ping()
		if err != nil {
			log.Fatal("Error: Could not establish a connection with the database")
		}
	}
	return db, err
}
