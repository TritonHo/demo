package handler

import (
	"errors"
	"net/http"

	"demo/lib/httputil"
	"demo/model"

	"github.com/go-xorm/xorm"
	"github.com/satori/go.uuid"
)

var errNotFound = errors.New("The record is not found.")

func CatGetOne(r *http.Request, urlValues map[string]string, session *xorm.Session) (int, error, interface{}) {
	//create the object and get the Id from the URL
	var cat model.Cat
	cat.Id = urlValues[`catId`]

	//load the object data from the database
	found, err := session.Id(cat.Id).Get(&cat)

	//output the object, or any error
	if err != nil {
		return http.StatusInternalServerError, err, nil
	} else {
		if found == false {
			return http.StatusNotFound, errNotFound, nil
		} else {
			return http.StatusOK, nil, cat
		}
	}
}

func CatGetAll(r *http.Request, urlValues map[string]string, session *xorm.Session) (int, error, interface{}) {
	//create the object slice
	cats := []model.Cat{}

	//load the object data from the database
	err := session.Find(&cats)

	if err != nil {
		return http.StatusInternalServerError, err, nil
	}

	//output the result
	return http.StatusOK, nil, cats
}

func CatUpdate(r *http.Request, urlValues map[string]string, session *xorm.Session) (int, error, interface{}) {
	id := urlValues[`catId`]

	//perform the input binding
	cat := model.Cat{}
	dbUpdateFields, _, err := httputil.BindForUpdate(r, &cat)
	//bind the input
	if err != nil {
		return http.StatusBadRequest, err, nil
	}

	//convert the columnName map into string slice
	columnNames := []string{}
	for k, _ := range dbUpdateFields {
		columnNames = append(columnNames, k)
	}

	//perform the update to the database
	affected, err := session.Where("id = ?", id).Cols(columnNames...).Update(&cat)

	//output the result
	if err != nil {
		return http.StatusInternalServerError, err, nil
	} else {
		if affected == 0 {
			return http.StatusNotFound, errNotFound, nil
		} else {
			return http.StatusNoContent, nil, nil
		}
	}
}

func CatCreate(r *http.Request, urlValues map[string]string, session *xorm.Session) (int, error, interface{}) {
	//bind the input
	cat := model.Cat{}
	if err := httputil.Bind(r, &cat); err != nil {
		return http.StatusBadRequest, err, nil
	}

	//generate the primary key for the cat
	cat.Id = uuid.NewV4().String()

	//perform the create to the database
	_, err := session.Insert(&cat)

	//output the result
	if err != nil {
		return http.StatusInternalServerError, err, nil
	} else {
		return http.StatusOK, nil, map[string]string{"Id": cat.Id}
	}
}

func CatDelete(r *http.Request, urlValues map[string]string, session *xorm.Session) (int, error, interface{}) {
	id := urlValues[`catId`]

	//perform the delete to the database
	affected, err := session.Id(id).Delete(new(model.Cat))

	//output the result
	if err != nil {
		return http.StatusInternalServerError, err, nil
	} else {
		if affected == 0 {
			return http.StatusNotFound, errNotFound, nil
		} else {
			return http.StatusNoContent, nil, nil
		}
	}
}
