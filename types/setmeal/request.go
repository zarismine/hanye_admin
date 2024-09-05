package setmeal

type SetmealDish struct {
	Copies    int     `json:"copies"`
	DishID    int     `json:"dishId"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	SetmealID int     `json:"setmealId"`
}

type AddRequest struct {
	CategoryId    int           `json:"categoryId"`
	Detail        string        `json:"detail"`
	Name          string        `json:"name"`
	Pic           string        `json:"pic"`
	Price         float64       `json:"price"`
	SetmealDishes []SetmealDish `json:"setmealDishes"`
	Status        int           `json:"status"`
}

type CategorySearchRequest struct {
	CategoryId int `json:"categoryId"`
}

type PageRequest struct {
	CategoryId int    `json:"categoryId"`
	Name       string `json:"name"`
	Page       int    `json:"page"`
	PageSize   int    `json:"pageSize"`
	Status     int    `json:"status"`
}

type EditRequest struct {
	Id int `json:"id"`
	AddRequest
}

type DeleteRequest struct {
	Ids []int `json:"ids"`
}
