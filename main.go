package main

import (
	"crypto/rsa"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"strconv"
	"time"

	"demo/handler"
	"demo/lib/auth"
	"demo/lib/config"
	"demo/lib/httputil"
	"demo/lib/middleware"
	"demo/setting"

	jwt "github.com/dgrijalva/jwt-go"
	xormCore "github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	initDependency()

	//in old go compiler, it is a must to enable multithread processing
	runtime.GOMAXPROCS(runtime.NumCPU())

	router := mux.NewRouter()
	uuidRegexp := `[[:alnum:]]{8}-[[:alnum:]]{4}-4[[:alnum:]]{3}-[89AaBb][[:alnum:]]{3}-[[:alnum:]]{12}`

	router.HandleFunc("/v1/auth/", middleware.Plain(handler.Login)).Methods("POST")

	//	router.HandleFunc("/v1/user", middleware.Plain(handler.UserCreate)).Methods("POST")
	//	router.HandleFunc("/v1/user", middleware.Plain(handler.UserUpdate)).Methods("PUT")

	router.HandleFunc("/v1/cats/", middleware.Wrap(handler.CatGetAll)).Methods("GET")
	router.HandleFunc("/v1/cats/{catId:"+uuidRegexp+"}", middleware.Wrap(handler.CatGetOne)).Methods("GET")
	router.HandleFunc("/v1/cats/{catId:"+uuidRegexp+"}", middleware.Wrap(handler.CatUpdate)).Methods("PUT")
	router.HandleFunc("/v1/cats/{catId:"+uuidRegexp+"}", middleware.Wrap(handler.CatDelete)).Methods("DELETE")
	router.HandleFunc("/v1/cats/", middleware.Wrap(handler.CatCreate)).Methods("POST")

	http.Handle("/", router)
	s := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	log.Fatal(s.ListenAndServe())
}

// init the various object and inject the database object to the modules
func initDependency() {
	//the postgresql connection string
	connectStr := "host=" + config.GetStr(setting.DB_HOST) +
		" port=" + strconv.Itoa(config.GetInt(setting.DB_PORT)) +
		" dbname=" + config.GetStr(setting.DB_NAME) +
		" user=" + config.GetStr(setting.DB_USERNAME) +
		" password='" + config.GetStr(setting.DB_PASSWORD) + "'" +
		" sslmode=disable"

	db, err := xorm.NewEngine("postgres", connectStr)
	if err != nil {
		log.Panic("DB connection initialization failed", err)
	}

	db.SetMaxIdleConns(config.GetInt(setting.DB_MAX_IDLE_CONN))
	db.SetMaxOpenConns(config.GetInt(setting.DB_MAX_OPEN_CONN))
	db.SetColumnMapper(xormCore.SnakeMapper{})
	//uncomment it if you want to debug
	//db.ShowSQL = true
	//db.ShowErr = true

	//load the RSA key from the file system, for the jwt auth
	var err1 error
	var currentKey *rsa.PrivateKey = nil
	var oldKey *rsa.PrivateKey = nil

	currentKeyBytes, _ := ioutil.ReadFile(config.GetStr(setting.JWT_RSA_KEY_LOCATION))
	currentKey, err1 = jwt.ParseRSAPrivateKeyFromPEM(currentKeyBytes)
	if err1 != nil {
		log.Panic(err1)
	}
	if location := config.GetStr(setting.JWT_OLD_RSA_KEY_LOCATION); location != `` {
		oldKeyBytes, _ := ioutil.ReadFile(location)
		oldKey, err1 = jwt.ParseRSAPrivateKeyFromPEM(oldKeyBytes)
		if err1 != nil {
			log.Panic(err1)
		}
	}
	lifetime := time.Duration(config.GetInt(setting.JWT_TOKEN_LIFETIME)) * time.Minute
	auth.Init(currentKey, oldKey, lifetime)

	httputil.Init(xormCore.SnakeMapper{})

	middleware.Init(db)
}
