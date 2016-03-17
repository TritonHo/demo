package handler

import (
	"net/http"

	"demo/lib/httputil"
	"demo/model"

	"github.com/go-xorm/xorm"
	"github.com/satori/go.uuid"
)

func CatGetOne(r *http.Request, urlValues map[string]string, session *xorm.Session, userId string) (int, error, interface{}) {
	//create the object and get the Id from the URL
	var cat model.Cat
	cat.Id = urlValues[`catId`]

	//load the object data from the database
	statusCode, err := getRecordWithUserId(&cat, cat.Id, userId, session)

	//output the object, or any error
	return statusCode, err, cat
}

func CatGetAll(r *http.Request, urlValues map[string]string, session *xorm.Session, userId string) (int, error, interface{}) {
	//create the object slice
	cats := []model.Cat{}

	//load the object data from the database
	err := session.Where(`user_id = ?`, userId).Find(&cats)

	if err != nil {
		return http.StatusInternalServerError, err, nil
	}

	//output the result
	return http.StatusOK, nil, cats
}

func CatUpdate(r *http.Request, urlValues map[string]string, session *xorm.Session, userId string) (int, error, interface{}) {
	id := urlValues[`catId`]

	//perform the input binding
	cat := model.Cat{}
	dbUpdateFields, _, err := httputil.BindForUpdate(r, &cat)
	//bind the input
	if err != nil {
		return http.StatusBadRequest, err, nil
	}

	//perform the update to the database
	statusCode, err := updateRecordWithUserId(&cat, dbUpdateFields, cat.Id, userId, session)

	//output the result
	return statusCode, err, nil
}

func CatCreate(r *http.Request, urlValues map[string]string, session *xorm.Session, userId string) (int, error, interface{}) {
	//bind the input
	cat := model.Cat{}
	if err := httputil.Bind(r, &cat); err != nil {
		return http.StatusBadRequest, err, nil
	}

	//generate the primary key for the cat
	cat.Id = uuid.NewV4().String()
	cat.UserId = userId

	//perform the create to the database
	statusCode, err := createRecord(&cat, session)

	//output the result
	if err != nil {
		return http.StatusInternalServerError, err, nil
	} else {
		return http.StatusOK, nil, map[string]string{"Id": cat.Id}
	}
}

func CatDelete(r *http.Request, urlValues map[string]string, session *xorm.Session, userId string) (int, error, interface{}) {
	id := urlValues[`catId`]

	//perform the delete to the database
	statusCode, err := deleteRecordWithUserId(new(model.Cat), id, userId, session)

	//output the result
	return statusCode, err, nil
}
