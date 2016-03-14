package main

import (
	"database/sql"
	"net/http"
	"runtime"
	//	"strconv"
	"time"

	"log"

	_ "github.com/lib/pq"
)

var (
	db *sql.DB
)

func main() {
	initObject()

	//in old go compiler, it is a must to enable multithread processing
	runtime.GOMAXPROCS(runtime.NumCPU())

	//http.HandleFunc(pattern string, handler func(ResponseWriter, *Request))
	http.HandleFunc(`/v1/cats/`, catHandler)

	s := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	log.Fatal(s.ListenAndServe())
}

// init the various object and inject the database object to the modules
func initObject() {
	//the postgresql connection string
	connectStr := "host=localhost" +
		" port=5432" +
		" dbname=demo_db" +
		" user=demo_user" +
		" password='user_password'" +
		" sslmode=disable"

	var err error = nil
	db, err = sql.Open("postgres", connectStr)
	if err != nil {
		log.Panic(err)
	}
}
