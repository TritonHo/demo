package model

import "time"

type Cat struct {
	Id string `xorm:"pk" json:"id"`

	Name   string `json:"name"`
	Gender string `json:"gender" validate:"required,enum=MALE/FEMALE"`

	CreateTime time.Time `xorm:"created" json:"createTime"`
	UpdateTime time.Time `xorm:"updated" json:"updateTime"`
}

func (c Cat) TableName() string {
	return "cats"
}
