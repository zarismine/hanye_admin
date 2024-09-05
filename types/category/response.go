package category

import "admin_app/model"

type AddResponse struct{}

type PageResponse struct {
	Records []*model.Category `json:"records"`
	Total   int               `json:"total"`
}
