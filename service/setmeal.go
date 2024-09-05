package service

import (
	"admin_app/global"
	"admin_app/model"
	"admin_app/model/jsontime"
	"admin_app/pkg/e"
	"admin_app/types/setmeal"
	"errors"
	"gorm.io/gorm"
	"time"
)

var SetmealService = newSetmealService()

func newSetmealService() *setmealService {
	return &setmealService{}
}

type setmealService struct {
}

func (cc *setmealService) AddSetmeal(req *setmeal.AddRequest, token string) int {
	createUser, _ := EmployeeService.GetEmployeeByToken(token)
	var count int64
	err := global.DB.Model(&model.Setmeal{}).Where("name = ?", req.Name).Count(&count).Error
	if count != 0 {
		return e.ERROR_ALREADY_EXIST
	}
	if err != nil {
		return e.ERROR_DATABASE
	}
	newSetmeal := &model.Setmeal{
		Name:       req.Name,
		Pic:        req.Pic,
		Detail:     req.Detail,
		Price:      req.Price,
		Status:     req.Status,
		CategoryId: req.CategoryId,
		CreateUser: createUser.Id,
		UpdateUser: createUser.Id,
		CreateTime: jsontime.JSONTime{Time: time.Now()},
		UpdateTime: jsontime.JSONTime{Time: time.Now()},
	}
	err = global.DB.Create(newSetmeal).Error
	if err != nil {
		return e.ERROR_DATABASE
	}
	for _, dish := range req.SetmealDishes {
		newDish := &model.SetmealDish{
			Name:      dish.Name,
			Price:     dish.Price,
			Copies:    dish.Copies,
			DishId:    dish.DishID,
			SetmealId: newSetmeal.Id,
		}
		err = global.DB.Create(newDish).Error
		if err != nil {
			return e.ERROR_DATABASE
		}
	}
	return e.SUCCESS
}

func (cc *setmealService) CategorySearchSetmealById(categoryId int) ([]*model.Setmeal, int) {
	var newSetmeals []*model.Setmeal
	err := global.DB.Table("setmeal").Where("category_id = ?", categoryId).Find(&newSetmeals).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, e.ERROR_DATABASE
	}
	return newSetmeals, e.SUCCESS
}

func (cc *setmealService) GetSetmealPage(req *setmeal.PageRequest) (*setmeal.PageResponse, int) {
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)
	var Setmeals []*model.Setmeal
	db := global.DB.Table("setmeal")
	if req.Name != "" {
		db = db.Where("name like ?", "%"+req.Name+"%")
	}
	if req.Status != -1 {
		db = db.Where("status = ?", req.Status)
	}
	if req.CategoryId != 0 {
		db = db.Where("category_id = ?", req.CategoryId)
	}
	err := db.Offset(offset).Limit(limit).Find(&Setmeals).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, e.ERROR_DATABASE
	}
	return &setmeal.PageResponse{
		Total:   len(Setmeals),
		Records: Setmeals,
	}, e.SUCCESS
}

func (cc *setmealService) GetSetmealById(id int) (*setmeal.SetmealAndDish, int) {
	var newSetmodel model.Setmeal
	err := global.DB.Table("setmeal").Where("id = ?", id).First(&newSetmodel).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, e.ERROR_EXIST
	}
	if err != nil {
		return nil, e.ERROR_DATABASE
	}
	var newSetmealDishes []*model.SetmealDish
	err = global.DB.Table("setmeal_dish").Where("setmeal_id = ?", id).Find(&newSetmealDishes).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, e.ERROR_DATABASE
	}
	return &setmeal.SetmealAndDish{
		Setmeal: newSetmodel,
		Dishes:  newSetmealDishes,
	}, e.SUCCESS
}

func (cc *setmealService) EditSetmealStatusById(id int, token string) int {
	editUser, _ := EmployeeService.GetEmployeeByToken(token)
	newSetmeal := new(model.Setmeal)
	err := global.DB.Where("id = ?", id).First(newSetmeal).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return e.ERROR_EXIST
	}
	if err != nil {
		return e.ERROR_DATABASE
	}
	var status int
	if newSetmeal.Status == 0 {
		status = 1
	}
	err = global.DB.Model(newSetmeal).Updates(map[string]interface{}{
		"status":      status,
		"update_user": editUser.Id,
	}).Error
	if err != nil {
		return e.ERROR_DATABASE
	}
	return e.SUCCESS
}

func (cc *setmealService) EditSetmeal(req *setmeal.EditRequest, token string) int {
	editUser, _ := EmployeeService.GetEmployeeByToken(token)
	newSetmeal := new(model.Setmeal)
	err := global.DB.Where("id = ?", req.Id).First(newSetmeal).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return e.INVALID_PARAMS
	}
	if err != nil {
		return e.ERROR_DATABASE
	}
	err = global.DB.Model(newSetmeal).Updates(map[string]interface{}{
		"name":        req.Name,
		"pic":         req.Pic,
		"detail":      req.Detail,
		"price":       req.Price,
		"status":      newSetmeal.Status,
		"category_id": req.CategoryId,
		"update_user": editUser.Id,
		"update_time": jsontime.JSONTime{Time: time.Now()},
	}).Error
	global.DB.Table("dish_flavor").Where("dish_id = ?", req.Id).Delete(&model.DishFlavor{})
	for _, SetmealDish := range req.SetmealDishes {
		newSetmealDish := &model.SetmealDish{
			Name:      SetmealDish.Name,
			Price:     SetmealDish.Price,
			Copies:    SetmealDish.Copies,
			DishId:    SetmealDish.DishID,
			SetmealId: newSetmeal.Id,
		}
		err = global.DB.Create(newSetmealDish).Error
		if err != nil {
			return e.ERROR_DATABASE
		}
	}
	return e.SUCCESS
}

func (cc *setmealService) DeleteSetmealByIds(req *setmeal.DeleteRequest) int {
	err := global.DB.Where("id in ?", req.Ids).Delete(&model.Setmeal{}).Error
	if err != nil {
		return e.ERROR_DATABASE
	}
	err = global.DB.Where("setmeal_id in ?", req.Ids).Delete(&model.SetmealDish{}).Error
	if err != nil {
		return e.ERROR_DATABASE
	}
	return e.SUCCESS
}
