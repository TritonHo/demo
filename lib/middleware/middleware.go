package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler func(r *http.Request, urlValues map[string]string) (statusCode int, err error, output interface{})

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
		if statusCode, err, output := f(req, mux.Vars(req)); err == nil {
			SendResponse(res, statusCode, output)
		} else {
			SendResponse(res, statusCode, map[string]string{"error": err.Error()})
		}
	}
}
