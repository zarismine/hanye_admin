package setmeal

import "admin_app/model"

type CategorySearchResponse struct{}

type PageResponse struct {
	Total   int              `json:"total"`
	Records []*model.Setmeal `json:"records"`
}

type SetmealAndDish struct {
	model.Setmeal
	Dishes []*model.SetmealDish
}
