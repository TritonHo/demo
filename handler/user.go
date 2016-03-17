package handler

import (
	//	"log"
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
