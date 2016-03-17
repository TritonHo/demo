package httputil

import (
	"bytes"
	"encoding/json"
	"net/http"
	"reflect"
	"regexp"
	"strings"

	xormCore "github.com/go-xorm/core"
)

var columnNameMapper xormCore.IMapper

func Init(mapper xormCore.IMapper) {
	columnNameMapper = mapper
}

// bind the http request to a struct. JSON, form, XML are supported
func Bind(r *http.Request, obj interface{}) error {
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(obj); err != nil {
		return err
	}

	return nil
}

//Only work for json/form input currently, not work for xml
func BindForUpdate(r *http.Request, obj interface{}) (dbFieldNames map[string]bool, fieldNames map[string]bool, e error) {
	keys := []string{}

	//FIXME: it may have security issue as it may use too much memory
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	inputBytes := buf.Bytes()

	if err := json.Unmarshal(inputBytes, obj); err != nil {
		return nil, nil, err
	}

	//get back the map
	t := map[string]interface{}{}
	if err := json.Unmarshal(inputBytes, &t); err != nil {
		return nil, nil, err
	} else {
		for k := range t {
			keys = append(keys, k)
		}
	}

	dbFieldNames, fieldNames = convertToFieldName(obj, keys)
	return dbFieldNames, fieldNames, nil
}

func GetXormColName(field *reflect.StructField) string {
	//a regular expression to trim the leading and tailing space
	re := regexp.MustCompile("^ *([a-zA-Z0-9-_]*) *")
	tag := field.Tag.Get(`xorm`)
	if tag == `-` {
		return ``
	}
	if tag != `` {
		for _, t := range strings.Split(tag, `,`) {
			//remove leading and ending space
			token := re.ReplaceAllString(t, "$1")

			//only take care the token with ''
			if strings.HasPrefix(token, `'`) && strings.HasSuffix(token, `'`) {
				return token[1 : len(token)-1]
			}
		}
	}

	return columnNameMapper.Obj2Table(field.Name)
}

//dbNameMap =  jsonName --> dbFieldName
//fieldNameMap =  jsonName --> fieldName
func getJsonTagMapping(obj interface{}) (dbNameMap map[string]string, fieldNameMap map[string]string) {
	//assumed a pointer will be passed
	immutable := reflect.ValueOf(obj).Elem()
	immutableType := immutable.Type()
	dbNameMap = map[string]string{}
	fieldNameMap = map[string]string{}

	for i := 0; i < immutable.NumField(); i++ {
		field := immutable.Field(i)
		fieldType := immutableType.Field(i)
		jsonName := getJsonTagName(&fieldType)

		if fieldType.Anonymous {
			//it is anonymous field, simply recursive dive in
			v := field.Addr().Interface()
			subDb, subField := getJsonTagMapping(v)
			for k, v := range subDb {
				dbNameMap[k] = v
			}
			for k, v := range subField {
				fieldNameMap[k] = v
			}
		} else {
			if field.CanSet() && jsonName != `` {
				dbNameMap[jsonName] = GetXormColName(&fieldType)
				fieldNameMap[jsonName] = fieldType.Name
			}
		}
	}
	return dbNameMap, fieldNameMap
}

//it accept the json field name, add use the structure json tag, to locate the structField name
func convertToFieldName(obj interface{}, jsonFieldName []string) (dbFieldNames map[string]bool, structFieldNames map[string]bool) {
	dbNameMap, fieldNameMap := getJsonTagMapping(obj)

	dbFieldsOutput := map[string]bool{}
	structFieldsOutput := map[string]bool{}
	for _, s := range jsonFieldName {
		if dbFieldName, ok := dbNameMap[s]; ok && dbFieldName != `` {
			dbFieldsOutput[dbFieldName] = true
		}
		if fieldName, ok := fieldNameMap[s]; ok {
			structFieldsOutput[fieldName] = true
		}
	}
	return dbFieldsOutput, structFieldsOutput
}

func getJsonTagName(field *reflect.StructField) string {
	if tag := field.Tag.Get(`json`); tag != `` {
		ss := strings.SplitN(tag, `,`, 2)
		if len(ss) >= 1 {
			return ss[0]
		}
	}
	return ``
}
