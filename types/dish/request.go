package dish

type Flavor struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	List string `json:"list"`
}

type AddRequest struct {
	Name       string   `json:"name"`
	Pic        string   `json:"pic"`
	Detail     string   `json:"detail"`
	Price      float64  `json:"price"`
	Status     int      `json:"status"`
	CategoryId int      `json:"categoryId"`
	Flavors    []Flavor `json:"flavors"`
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
