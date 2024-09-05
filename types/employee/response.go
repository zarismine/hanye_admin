package employee

import "admin_app/model"

type RegisterResponse struct {
}

type LoginResponse struct {
	Id      int    `json:"id"`
	Token   string `json:"token"`
	Account string `json:"account"`
}
type EditPasswordResponse struct {
}
type AddEditResponse struct {
}

type PageResponse struct {
	Total   int               `json:"total"`
	Records []*model.Employee `json:"records"`
}
