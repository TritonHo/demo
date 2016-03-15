package validate

//FIXME: the "required" tag doesn't fullfill the zero value of the int/float

//for details in this library, please reference to http://godoc.org/gopkg.in/bluesuncorp/validator.v8

import (
	"errors"
	"reflect"
	"strings"

	valid "gopkg.in/bluesuncorp/validator.v8"
)

const (
	ValidatorTag string = "validate"
	//copied from "gopkg.in/bluesuncorp/validator.v8", DO NOT CHANGE IT
	TagSeparator    string = ","
	OrSeparator     string = "|"
	TagKeySeparator string = "="
)

//TODO: add rules to check UUID datatype
var createValidator *valid.Validate
var updateValidator *valid.Validate

func init() {
	createValidator = valid.New(&valid.Config{TagName: ValidatorTag})
	updateValidator = valid.New(&valid.Config{TagName: ValidatorTag})

	createValidator.RegisterValidation("enum", enum)
	createValidator.RegisterValidation("fixed", fake)
	createValidator.RegisterValidation("zerotime", zerotime)
	updateValidator.RegisterValidation("enum", enum)
	updateValidator.RegisterValidation("fixed", fixed)
	updateValidator.RegisterValidation("zerotime", zerotime)
}

func zerotime(v *valid.Validate, topStruct reflect.Value, currentStructOrField reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	//always return true, as the zerotime will be handled by the httputil
	return true
}

//this procedure implement the "enum" validation
//for example, Cat.Gender,  should be one of value in [MALE, FEMALE]
//thus, the structTag will be validate:"MALE/FEMALE
//remarks: this function use "/" as a delimitor between values, and no comma / space is allowed in the value
func enum(v *valid.Validate, topStruct reflect.Value, currentStructOrField reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {

	if fieldKind != reflect.String {
		return false
	}
	s := field.String()
	for _, candidate := range strings.Split(param, `/`) {
		if s == candidate {
			return true
		}
	}
	return s == `` //empty string is also no problem
}

func fake(v *valid.Validate, topStruct reflect.Value, currentStructOrField reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	return true
}

func fixed(v *valid.Validate, topStruct reflect.Value, currentStructOrField reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	//fixed attribute should not in the input in the PUT method
	return false
}

//s must be struct only, other type is not allowed
func ValidateStructForCreate(s interface{}) error {
	if err := createValidator.Struct(s); err != nil {
		return err
	}
	return nil
}

//s must be struct only, other type is not allowed
func ValidateStructForUpdate(s interface{}, structFieldNames map[string]bool) error {
	if len(structFieldNames) == 0 {
		return errors.New("No attribute is provided")
	}

	//TODO: add validation on the primary key, the pk should not be updatable

	immutable := reflect.ValueOf(s).Elem()
	immutableType := immutable.Type()
	for fieldName, _ := range structFieldNames {
		field := immutable.FieldByName(fieldName)
		fieldType, _ := immutableType.FieldByName(fieldName)

		if err := updateValidator.Field(field.Interface(), fieldType.Tag.Get(ValidatorTag)); err != nil {
			return err
		}
	}

	return nil
}
