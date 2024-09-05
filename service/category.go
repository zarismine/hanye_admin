package service

import (
	"admin_app/global"
	"admin_app/model"
	"admin_app/model/jsontime"
	"admin_app/pkg/e"
	"admin_app/types/category"
	"errors"
	"gorm.io/gorm"
	"time"
)

var CategoryService = newCategoryService()

func newCategoryService() *categoryService {
	return &categoryService{}
}

type categoryService struct {
}

func (cc *categoryService) AddCategory(req *category.AddRequest, token string) int {
	// 参数校验省略
	if req.Id != 0 {
		return e.INVALID_PARAMS
	}
	var t int64
	global.DB.Table("category").Where("name = ?", req.Name).Count(&t)
	if t != 0 {
		return e.ERROR_ALREADY_EXIST
	}
	createUser, errcode := EmployeeService.GetEmployeeByToken(token)
	if errcode != e.SUCCESS {
		return errcode
	}
	newCategory := &model.Category{
		Name:       req.Name,
		Type:       req.Type,
		Sort:       req.Sort,
		CreateUser: createUser.Id,
		UpdateUser: createUser.Id,
		CreateTime: jsontime.JSONTime{Time: time.Now()},
		UpdateTime: jsontime.JSONTime{Time: time.Now()},
	}
	err := global.DB.Create(newCategory).Error
	if err != nil {
		return e.ERROR_DATABASE
	}
	return e.SUCCESS
}

func (cc *categoryService) PageCategory(req *category.PageRequest) (*category.PageResponse, int) {
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)
	var Categories []*model.Category
	db := global.DB.Table("category")
	if req.Name != "" {
		db = db.Where("name like ?", "%"+req.Name+"%")
	}

	if req.Type != 0 {
		db = db.Where("type = ?", req.Type)
	}
	err := db.Offset(offset).Limit(limit).Find(&Categories).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, e.ERROR_DATABASE
	}
	return &category.PageResponse{
		Total:   len(Categories),
		Records: Categories,
	}, e.SUCCESS
}

func (cc *categoryService) GetDetailById(id int) (*model.Category, int) {
	newCategory := new(model.Category)
	err := global.DB.Where("id = ?", id).First(newCategory).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, e.INVALID_PARAMS
	}
	if err != nil {
		return nil, e.ERROR_DATABASE
	}
	return newCategory, e.SUCCESS
}

func (cc *categoryService) EditCategoryStatus(id int, token string) int {
	OptUser, errcode := EmployeeService.GetEmployeeByToken(token)
	if errcode != e.SUCCESS {
		return errcode
	}
	newCategory := new(model.Category)
	err := global.DB.Where("id = ?", id).First(newCategory).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return e.INVALID_PARAMS
	}
	if err != nil {
		return e.ERROR_DATABASE
	}
	var status int
	if newCategory.Status == 0 {
		status = 1
	}
	err = global.DB.Model(newCategory).Updates(map[string]interface{}{
		"status":      status,
		"update_user": OptUser.Id,
	}).Error
	if err != nil {
		return e.ERROR_DATABASE
	}
	return e.SUCCESS
}

func (cc *categoryService) EditCategory(req *category.EditRequest, token string) int {
	OptUser, errcode := EmployeeService.GetEmployeeByToken(token)
	if errcode != e.SUCCESS {
		return errcode
	}
	newCategory := new(model.Category)
	err := global.DB.Where("id = ?", req.Id).First(newCategory).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return e.INVALID_PARAMS
	}
	if err != nil {
		return e.ERROR_DATABASE
	}
	err = global.DB.Model(newCategory).Updates(map[string]interface{}{
		"name":        req.Name,
		"type":        req.Type,
		"sort":        req.Sort,
		"update_user": OptUser.Id,
	}).Error
	if err != nil {
		return e.ERROR_DATABASE
	}
	return e.SUCCESS
}

func (cc *categoryService) DeleteCategoryById(id int) int {
	newCategory := new(model.Category)
	err := global.DB.Where("id = ?", id).First(newCategory).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return e.INVALID_PARAMS
	}
	if err != nil {
		return e.ERROR_DATABASE
	}
	err = global.DB.Delete(newCategory).Error
	if err != nil {
		return e.ERROR_DATABASE
	}
	return e.SUCCESS
}

func (cc *categoryService) GetListByType(req *category.GetListRequest) ([]*model.Category, int) {
	typeReq := req.Type
	var categories []*model.Category
	err := global.DB.Where("type = ", typeReq).Find(&categories).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, e.INVALID_PARAMS
	}
	if err != nil {
		return nil, e.ERROR_DATABASE
	}
	return categories, e.SUCCESS
}
