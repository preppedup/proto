package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var db *sql.DB

func init(){
	var err error
	db, err = sql.Open("postgres", "user=devadmin password=devadmin123 port=5432 host=dev-db.cn2axrsupztc.ap-southeast-1.rds.amazonaws.com dbname=preppedupdb sslmode=disable")
	if err != nil {
		log.Fatal("Error: Could not establish a connection with the database" + err.Error())
	}

	err2 := db.Ping()
	if err2 != nil {
		log.Fatal("Error: Could not ping the database:" + err2.Error())
	} else{
		log.Print("Connected with db.")
	}

}

func Pool() (*sql.DB) {
	if (db == nil){
		log.Print("Error: We lost the db! This might be a good time to panic.")
	}
	return db
}
