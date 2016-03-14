package handler

import (
	"github.com/go-xorm/xorm"
)

var (
	db *xorm.Engine
)

func Init(database *xorm.Engine) {
	db = database
}
