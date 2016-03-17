package handler

import (
	"errors"
	"net/http"

	"demo/lib/auth"
	"demo/lib/httputil"
	"demo/lib/middleware"
	"demo/model"

	"github.com/go-xorm/xorm"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

func UserCreate(w http.ResponseWriter, r *http.Request, urlValues map[string]string, db *xorm.Engine) {
	user := struct {
		model.User `xorm:"extends"`
		Password   string `xorm:"-" json:"password" validate:"required"`
	}{}

	if err := httputil.Bind(r, &user); err != nil {
		middleware.SendResponse(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	user.Id = uuid.NewV4().String()

	if digest, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost); err != nil {
		middleware.SendResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	} else {
		user.PasswordDigest = string(digest)
	}

	q := `	insert into users(id, email, password_digest, first_name, last_name)
			select ?, ?, ?, ?, ? 
			where not exists (select 1 from users where email = ?)`

	result, err := db.Exec(q, user.Id, user.Email, user.PasswordDigest, user.FirstName, user.LastName, user.Email)
	if err != nil {
		middleware.SendResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	if affected, _ := result.RowsAffected(); affected == 0 {
		middleware.SendResponse(w, http.StatusForbidden, map[string]string{"error": "The email is already used."})
		return
	}

	if newToken, err := auth.Sign(user.Id); err != nil {
		middleware.SendResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
	} else {
		// update JWT Token
		w.Header().Add("Authorization", newToken)
		//allow CORS
		w.Header().Set("Access-Control-Expose-Headers", "Authorization")
		middleware.SendResponse(w, http.StatusOK, map[string]string{"userId": user.Id})
	}
}

func UserUpdate(r *http.Request, urlValues map[string]string, session *xorm.Session, userId string) (int, error, interface{}) {
	id := urlValues[`userId`]
	if id != userId {
		return http.StatusForbidden, errors.New("Updating others account is forbidden"), nil
	}

	input := struct {
		model.User       `xorm:"extends"`
		Password         *string `xorm:"-" json:"password" validate:"omitempty,min=1"`
		OriginalPassword *string `xorm:"-" json:"originalPassword" validate:"omitempty,min=1"`
	}{}

	//perform the input binding
	dbUpdateFields, _, err := httputil.BindForUpdate(r, &input)
	//bind the input
	if err != nil {
		return http.StatusBadRequest, err, nil
	}
	if input.Password != nil {
		//if user changes the password, he must provide the original password
		if input.OriginalPassword == nil {
			return http.StatusForbidden, errors.New("Please provide the original password"), nil
		} else {
			//get the user record from database, and ensure the original password is correct
			user := model.User{}
			if found, err := session.Id(userId).Get(&user); !found {
				return http.StatusNotFound, errNotFound, nil
			} else {
				if err != nil {
					return http.StatusInternalServerError, err, nil
				}
				if bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(*input.OriginalPassword)) != nil {
					return http.StatusForbidden, errors.New(`The original password is invalid`), nil
				}
			}

			//generate the bcrypt hash with the new password
			if digest, err := bcrypt.GenerateFromPassword([]byte(*input.Password), bcrypt.DefaultCost); err != nil {
				return http.StatusInternalServerError, err, nil
			} else {
				input.PasswordDigest = string(digest)
				dbUpdateFields[`password_digest`] = true
			}
		}
	}

	//convert the columnName map into string slice
	columnNames := []string{}
	for k, _ := range dbUpdateFields {
		columnNames = append(columnNames, k)
	}

	//perform the update to the database
	affected, err := session.Id(userId).Cols(columnNames...).Update(&input)

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
