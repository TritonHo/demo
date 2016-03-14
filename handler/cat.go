package handler

import (
	"encoding/json"
	"net/http"

	"demo/lib/httputil"
	"demo/model"

	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
)

func CatGetOne(w http.ResponseWriter, r *http.Request) {
	//create the object and get the Id from the URL
	var cat model.Cat
	cat.Id = mux.Vars(r)[`catId`]

	//load the object data from the database
	found, err := db.Id(cat.Id).Get(&cat)

	//perform the object, or any error
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error":"` + err.Error() + `"}`))
	} else {
		if found == false {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(cat)
		}
	}
}

func CatGetAll(w http.ResponseWriter, r *http.Request) {
	//create the object slice
	cats := []model.Cat{}

	//load the object data from the database
	err := db.Find(&cats)

	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error":"` + err.Error() + `"}`))
		return
	}

	//output the result
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cats)
}

func CatUpdate(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)[`catId`]

	//perform the input binding
	cat := model.Cat{}
	dbUpdateFields, updateFields, err := httputil.BindForUpdate(r, &cat)
	//bind the input
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":"` + err.Error() + `"}`))
		return
	}
	//perform basic checking on gender
	if _, ok := updateFields[`Gender`]; ok && cat.Gender != `MALE` && cat.Gender != `FEMALE` {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":"Gender must be MALE or FEMALE"}`))
		return
	}

	//convert the columnName map into string slice
	columnNames := []string{}
	for k, _ := range dbUpdateFields {
		columnNames = append(columnNames, k)
	}

	//perform the update to the database
	affected, err := db.Where("id = ?", id).Cols(columnNames...).Update(&cat)

	//output the result
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error":"` + err.Error() + `"}`))
	} else {
		if affected == 0 {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusNoContent)
		}
	}
}

func CatCreate(w http.ResponseWriter, r *http.Request) {
	//bind the input
	cat := model.Cat{}
	if err := httputil.Bind(r, &cat); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":"` + err.Error() + `"}`))
		return
	}
	//perform basic checking on gender
	if cat.Gender != `MALE` && cat.Gender != `FEMALE` {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":"Gender must be MALE or FEMALE"}`))
		return
	}

	//generate the primary key for the cat
	cat.Id = uuid.NewV4().String()

	//perform the create to the database
	_, err := db.Insert(&cat)

	//output the result
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error":"` + err.Error() + `"}`))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"id":"` + cat.Id + `"}`))
	}
}

func CatDelete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)[`catId`]

	//perform the delete to the database
	affected, err := db.Id(id).Insert(new(model.Cat))

	//output the result
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error":"` + err.Error() + `"}`))
	} else {
		if affected == 0 {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusNoContent)
		}
	}
}
