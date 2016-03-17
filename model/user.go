package model

import "time"

type User struct {
	Id string `xorm:"pk" json:"id" validate:"fixed"`

	Email          string `json:"email" validate:"fixed"`
	PasswordDigest string `json:"-"`

	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`

	CreateTime time.Time `xorm:"created" json:"createTime" validate:"zerotime"`
	UpdateTime time.Time `xorm:"updated" json:"updateTime" validate:"zerotime"`
}

func (c User) TableName() string {
	return "users"
}
