package model

import (
	"admin_app/model/jsontime"
)

type Category struct {
	Id         int               `gorm:"column:id;type:int;primaryKey;" json:"id"`
	Name       string            `gorm:"column:name;type:varchar(64);not null;" json:"name"`
	Type       int               `gorm:"column:type;type:tinyint;not null;" json:"type"`
	Sort       int               `gorm:"column:sort;type:int;not null;" json:"sort"`
	Status     int               `gorm:"column:status;type:tinyint;not null;default:1;" json:"status"`
	CreateUser int               `gorm:"column:create_user;type:int;not null;" json:"createUser"`
	UpdateUser int               `gorm:"column:update_user;type:int;not null;" json:"updateUser"`
	CreateTime jsontime.JSONTime `gorm:"column:create_time;type:datetime;not null;" json:"createTime"`
	UpdateTime jsontime.JSONTime `gorm:"column:update_time;type:datetime;not null;" json:"updateTime"`
}

func (Category) TableName() string {
	return "category"
}
