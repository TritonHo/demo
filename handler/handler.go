package handler

import (
	"errors"
	"net/http"

	"github.com/go-xorm/xorm"
)

var (
	errNotFound = errors.New("The record is not found.")
)

//the id should be a uuid
func getRecord(out interface{}, id string, session *xorm.Session) (statusCode int, err error) {
	found, err := session.Id(id).Get(out)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if found == false {
		return http.StatusNotFound, errNotFound
	}

	return http.StatusOK, nil
}

func getRecordWithUserId(out interface{}, id, userId string, session *xorm.Session) (statusCode int, err error) {
	found, err := session.Where("id = ? and user_id = ?", id, userId).Get(out)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if found == false {
		return http.StatusNotFound, errNotFound
	}

	return http.StatusOK, nil
}

func deleteRecordWithUserId(input interface{}, id, userId string, session *xorm.Session) (statusCode int, err error) {
	affectedCount, err := session.Where("id = ? and user_id = ?", id, userId).Delete(input)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if affectedCount == 0 {
		return http.StatusNotFound, errNotFound
	}

	return http.StatusNoContent, nil
}

func updateRecordWithUserId(input interface{}, fieldNames map[string]bool, id, userId string, session *xorm.Session) (statusCode int, err error) {
	//convert the fields set to array
	array := []string{}
	for k, _ := range fieldNames {
		array = append(array, k)
	}

	//update the database
	affected, err := session.Where("id = ? and user_id = ?", id, userId).Cols(array...).Update(input)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if affected == 0 {
		return http.StatusNotFound, errNotFound
	}

	return http.StatusNoContent, nil
}

func createRecord(input interface{}, session *xorm.Session) (statusCode int, err error) {
	_, err = session.Insert(input)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, err
}
