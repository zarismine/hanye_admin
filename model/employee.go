package model

import (
	"admin_app/model/jsontime"
)

type Employee struct {
	Id         int               `json:"id" form:"id" gorm:"column:id"`
	Name       string            `json:"name" form:"name" gorm:"column:name"`
	Account    string            `json:"account" form:"account" gorm:"column:account"`
	Password   string            `json:"password" form:"password" gorm:"column:password"`
	Phone      string            `json:"phone" form:"phone" gorm:"column:phone"`
	Age        int               `json:"age" form:"age" gorm:"column:age"`
	Gender     int               `json:"gender" form:"gender" gorm:"column:gender"`
	Pic        string            `json:"pic" form:"pic" gorm:"column:pic"`
	Status     int               `json:"status" form:"status" gorm:"column:status"`
	CreateUser int               `json:"createUser" form:"createUser" gorm:"column:create_user"`
	UpdateUser int               `json:"updateUser" form:"updateUser" gorm:"column:update_user"`
	CreateTime jsontime.JSONTime `json:"createTime" form:"createTime" gorm:"column:create_time;type:datetime"`
	UpdateTime jsontime.JSONTime `json:"updateTime" form:"updateTime" gorm:"column:update_time;type:datetime"`
}

func (Employee) TableName() string {
	return "employee"
}
