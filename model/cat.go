package model

import "time"

type Cat struct {
	Id     string `xorm:"pk" json:"id"`
	UserId string `json:"UserId" validate:"fixed"`

	Name   string `json:"name"`
	Gender string `json:"gender" validate:"required,enum=MALE/FEMALE"`

	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
}

func (c Cat) TableName() string {
	return "cats"
}
