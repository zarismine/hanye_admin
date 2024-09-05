package model

type SetmealDish struct {
	Id        int     `gorm:"column:id;type:int;primaryKey;" json:"id"`
	Name      string  `gorm:"column:name;type:varchar(64);not null;" json:"name"`
	Price     float64 `gorm:"column:price;type:decimal(10, 2);not null;" json:"price"`
	Copies    int     `gorm:"column:copies;type:int;not null;" json:"copies"`
	DishId    int     `gorm:"column:dish_id;type:int;not null;" json:"dish_id"`
	SetmealId int     `gorm:"column:setmeal_id;type:int;not null;" json:"setmeal_id"`
}

func (SetmealDish) TableName() string {
	return "setmeal_dish"
}
