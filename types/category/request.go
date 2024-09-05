package category

type AddRequest struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Type int    `json:"type"`
	Sort int    `json:"sort"`
}

type PageRequest struct {
	Name     string `json:"name"`
	Page     int    `json:"page"`
	PageSize int    `json:"pageSize"`
	Type     int    `json:"type"`
}

type EditRequest struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Type int    `json:"type"`
	Sort int    `json:"sort"`
}

type GetListRequest struct {
	Type int `json:"type"`
}
