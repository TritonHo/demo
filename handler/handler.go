package handler

import (
	"errors"
	"net/http"
	"reflect"
	"strconv"

	"demo/lib/config"
	"demo/setting"

	"github.com/go-xorm/xorm"
)

var (
	errNotFound = errors.New("The record is not found.")
)

type PagingOutput struct {
	//the number of record that fulfill the searching criteria
	Total int `json:"total"`
	//the starting position of the first record in the output
	//the index is in zero based
	StartingIndex int         `json:"startingIndex"`
	Data          interface{} `json:"data"`
}

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
	found, err := session.Where("id = ? and user_id = ?", id, userId).Desc(`create_time`).Get(out)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if found == false {
		return http.StatusNotFound, errNotFound
	}

	return http.StatusOK, nil
}

func getPagingInput(r *http.Request) (startIndex int, limit int) {
	startIndex = 0
	if t, err := strconv.Atoi(r.FormValue("start")); err == nil {
		startIndex = t
	}
	limit = config.GetInt(setting.DEFAULT_RECORD_PER_PAGE)
	if t, err := strconv.Atoi(r.FormValue("limit")); err == nil {
		limit = t
	}
	return
}
func getAllRecordWithUserId(input interface{}, startIndex int, limit int, userId string, session *xorm.Session) (statusCode int, err error, output *PagingOutput) {
	//create the slice pointer by reflection
	immutable := reflect.ValueOf(input).Elem()
	sliceType := reflect.SliceOf(immutable.Type())
	slice := reflect.New(sliceType).Interface()

	err0 := session.Where("user_id = ?", userId).Limit(limit, startIndex).Asc(`create_time`).Find(slice)
	if err0 != nil {
		return http.StatusInternalServerError, err0, nil
	}

	total, err1 := session.Where("user_id = ?", userId).Count(input)
	if err1 != nil {
		return http.StatusInternalServerError, err1, nil
	}

	//output an empty json array if no record found
	if reflect.ValueOf(slice).Elem().Len() == 0 {
		slice = []interface{}{}
	}

	return http.StatusOK, nil, &PagingOutput{Total: int(total), StartingIndex: startIndex, Data: slice}
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
