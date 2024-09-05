package model

type DishFlavor struct {
	Id     int    `gorm:"column:id;type:int;primaryKey;" json:"id"`
	Name   string `gorm:"column:name;type:varchar(64);not null;" json:"name"`
	List   string `gorm:"column:list;type:varchar(255);not null;" json:"list"`
	DishId int    `gorm:"column:dish_id;type:int;not null;" json:"dishId"`
}

func (DishFlavor) TableName() string {
	return "dish_flavor"
}
