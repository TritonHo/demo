package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/go-xorm/xorm"
	"github.com/gorilla/mux"
)

var (
	db *xorm.Engine
)

func Init(database *xorm.Engine) {
	db = database
}

type Handler func(r *http.Request, urlValues map[string]string, session *xorm.Session) (statusCode int, err error, output interface{})

//type PlainHandler func(res http.ResponseWriter, req *http.Request, urlValues map[string]string)

// send a http response to the user with JSON format
func SendResponse(res http.ResponseWriter, statusCode int, data interface{}) {
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.WriteHeader(statusCode)
	if d, ok := data.([]byte); ok {
		res.Write(d)
	} else {
		json.NewEncoder(res).Encode(data)
	}
}

// a middleware to handle user authorization
func Wrap(f Handler) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		//prepare a database session for the handler
		session := db.NewSession()
		if err := session.Begin(); err != nil {
			SendResponse(res, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
		defer session.Close()

		//everything seems fine, goto the business logic handler
		if statusCode, err, output := f(req, mux.Vars(req), session); err == nil {
			//the business logic handler return no error, then try to commit the db session
			if err := session.Commit(); err != nil {
				SendResponse(res, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			} else {
				SendResponse(res, statusCode, output)
			}
		} else {
			session.Rollback()
			SendResponse(res, statusCode, map[string]string{"error": err.Error()})
		}
	}
}
