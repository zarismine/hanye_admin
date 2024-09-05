package service

import (
	"admin_app/global"
	"admin_app/model"
	"admin_app/model/jsontime"
	"admin_app/pkg/e"
	"admin_app/types/dish"
	"errors"
	"gorm.io/gorm"
	"time"
)

var DishService = newDishService()

func newDishService() *dishService {
	return &dishService{}
}

type dishService struct {
}

func (cc *dishService) AddDish(req *dish.AddRequest, token string) int {
	createUser, _ := EmployeeService.GetEmployeeByToken(token)
	var count int64
	err := global.DB.Model(&model.Dish{}).Where("name = ?", req.Name).Count(&count).Error
	if count != 0 {
		return e.ERROR_ALREADY_EXIST
	}
	if err != nil {
		return e.ERROR_DATABASE
	}
	newDish := &model.Dish{
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
	err = global.DB.Create(newDish).Error
	if err != nil {
		return e.ERROR_DATABASE
	}
	for _, flavor := range req.Flavors {
		newFlavor := &model.DishFlavor{
			Name:   flavor.Name,
			List:   flavor.List,
			DishId: newDish.Id,
		}
		err = global.DB.Create(newFlavor).Error
		if err != nil {
			return e.ERROR_DATABASE
		}
	}
	return e.SUCCESS
}

func (cc *dishService) GetDishPage(req *dish.PageRequest) (*dish.PageResponse, int) {
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)
	var dishes []*dish.Dish
	db := global.DB.Table("dish")
	if req.Name != "" {
		db = db.Where("name like ?", "%"+req.Name+"%")
	}
	if req.Status != -1 {
		db = db.Where("status = ?", req.Status)
	}
	if req.CategoryId != 0 {
		db = db.Where("category_id = ?", req.CategoryId)
	}
	err := db.Offset(offset).Limit(limit).Find(&dishes).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, e.ERROR_DATABASE
	}
	return &dish.PageResponse{
		Total:   len(dishes),
		Records: dishes,
	}, e.SUCCESS
}

func (cc *dishService) GetDishAndFlavorBy(id int) (*dish.FlavorDishResponse, int) {
	var newDish dish.Dish
	err := global.DB.Table("dish").Where("id = ?", id).First(&newDish).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, e.ERROR_EXIST
	}
	if err != nil {
		return nil, e.ERROR_DATABASE
	}
	var newFlavors []*model.DishFlavor
	err = global.DB.Table("dish_flavor").Where("dish_id = ?", id).Find(&newFlavors).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, e.ERROR_DATABASE
	}
	return &dish.FlavorDishResponse{
		Dish:    newDish,
		Flavors: newFlavors,
	}, e.SUCCESS
}

func (cc *dishService) EditDishStatusById(id int, token string) int {
	editUser, _ := EmployeeService.GetEmployeeByToken(token)
	newDish := new(model.Dish)
	err := global.DB.Where("id = ?", id).First(newDish).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return e.ERROR_EXIST
	}
	if err != nil {
		return e.ERROR_DATABASE
	}
	var status int
	if newDish.Status == 0 {
		status = 1
	}
	err = global.DB.Model(newDish).Updates(map[string]interface{}{
		"status":      status,
		"update_user": editUser.Id,
	}).Error
	if err != nil {
		return e.ERROR_DATABASE
	}
	return e.SUCCESS
}

func (cc *dishService) EditDish(req *dish.EditRequest, token string) int {
	editUser, _ := EmployeeService.GetEmployeeByToken(token)
	newDish := new(model.Dish)
	err := global.DB.Where("id = ?", req.Id).First(newDish).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return e.INVALID_PARAMS
	}
	if err != nil {
		return e.ERROR_DATABASE
	}
	err = global.DB.Model(newDish).Updates(map[string]interface{}{
		"name":        req.Name,
		"pic":         req.Pic,
		"detail":      req.Detail,
		"price":       req.Price,
		"status":      newDish.Status,
		"category_id": req.CategoryId,
		"update_user": editUser.Id,
		"update_time": jsontime.JSONTime{Time: time.Now()},
	}).Error
	global.DB.Table("dish_flavor").Where("dish_id = ?", req.Id).Delete(&model.DishFlavor{})
	for _, flavor := range req.Flavors {
		newFlavor := &model.DishFlavor{
			Name:   flavor.Name,
			List:   flavor.List,
			DishId: newDish.Id,
		}
		err = global.DB.Create(newFlavor).Error
		if err != nil {
			return e.ERROR_DATABASE
		}
	}
	return e.SUCCESS
}

func (cc *dishService) DeleteDishByIds(req *dish.DeleteRequest) int {
	err := global.DB.Where("id in ?", req.Ids).Delete(&model.Dish{}).Error
	if err != nil {
		return e.ERROR_DATABASE
	}
	err = global.DB.Where("dish_id in ?", req.Ids).Delete(&model.DishFlavor{}).Error
	if err != nil {
		return e.ERROR_DATABASE
	}
	return e.SUCCESS
}
