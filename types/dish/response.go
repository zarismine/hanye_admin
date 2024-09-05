package dish

import (
	"admin_app/model"
	"admin_app/model/jsontime"
)

type Dish struct {
	Id         int               `json:"id"`
	Name       string            `json:"name"`
	Pic        string            `json:"pic"`
	Detail     string            `json:"detail"`
	Price      float64           `json:"price"`
	Status     int               `json:"status"`
	CategoryId int               `json:"categoryId"`
	UpdateTime jsontime.JSONTime `json:"updateTime"`
}

type PageResponse struct {
	Total   int     `json:"total"`
	Records []*Dish `json:"records"`
}

type FlavorDishResponse struct {
	Dish
	Flavors []*model.DishFlavor `json:"flavors"`
}
