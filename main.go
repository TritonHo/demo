package main

import (
	"database/sql"
	"net/http"
	"runtime"
	"time"

	"log"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var (
	db *sql.DB
)

func main() {
	initObject()

	//in old go compiler, it is a must to enable multithread processing
	runtime.GOMAXPROCS(runtime.NumCPU())

	router := mux.NewRouter()
	uuidRegexp := `[[:alnum:]]{8}-[[:alnum:]]{4}-4[[:alnum:]]{3}-[89AaBb][[:alnum:]]{3}-[[:alnum:]]{12}`

	router.HandleFunc("/v1/cats/", catGetAll).Methods("GET")
	router.HandleFunc("/v1/cats/{catId:"+uuidRegexp+"}", catGetOne).Methods("GET")
	router.HandleFunc("/v1/cats/{catId:"+uuidRegexp+"}", catUpdate).Methods("PUT")
	router.HandleFunc("/v1/cats/{catId:"+uuidRegexp+"}", catDelete).Methods("DELETE")
	router.HandleFunc("/v1/cats/", catCreate).Methods("POST")

	http.Handle("/", router)
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
