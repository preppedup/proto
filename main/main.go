package main

import (
	"fmt"
	"log"
	"net/http"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/gorilla/mux"
)

func main() {

	db, err := sql.Open("postgres", "user=devadmin password=devadmin123 port=5432 host=dev-db.cn2axrsupztc.ap-southeast-1.rds.amazonaws.com dbname=dev-db  sslmode=disable")
	if err != nil {
		log.Fatal("Error: Could not establish a connection with the database")
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Error: Could not establish a connection with the database")
	}

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/mobiles/{mobileNum}/otps",OtpGen)
	router.HandleFunc("/mobiles/{mobileNum}/otps/{otp}",OtpVerify)

	log.Fatal(http.ListenAndServe(":8080",router))



}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")


}

func OtpGen(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Otp generation not live yet!")
}

func OtpVerify(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	otp := vars["otp"]
	fmt.Fprintln(w, "You tried to verify ", otp)
}
