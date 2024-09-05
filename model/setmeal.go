package model

import (
	"admin_app/model/jsontime"
)

type Setmeal struct {
	Id         int               `gorm:"column:id;type:int;primaryKey;" json:"id"`
	Name       string            `gorm:"column:name;type:varchar(64);not null;" json:"name"`
	Pic        string            `gorm:"column:pic;type:longtext;" json:"pic"`
	Detail     string            `gorm:"column:detail;type:varchar(255);not null;" json:"detail"`
	Price      float64           `gorm:"column:price;type:decimal(10, 2);not null;" json:"price"`
	Status     int               `gorm:"column:status;type:tinyint;not null;default:1;" json:"status"`
	CategoryId int               `gorm:"column:category_id;type:int;not null;" json:"categoryId"`
	CreateUser int               `gorm:"column:create_user;type:int;not null;" json:"createUser"`
	UpdateUser int               `gorm:"column:update_user;type:int;not null;" json:"updateUser"`
	CreateTime jsontime.JSONTime `gorm:"column:create_time;type:datetime;" json:"createTime"`
	UpdateTime jsontime.JSONTime `gorm:"column:update_time;type:datetime;" json:"updateTime"`
}

func (Setmeal) TableName() string {
	return "setmeal"
}
