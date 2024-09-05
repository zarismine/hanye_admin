package service

import (
	"admin_app/global"
	"admin_app/model"
	"admin_app/model/jsontime"
	"admin_app/pkg/e"
	"admin_app/pkg/utils"
	"admin_app/types/employee"
	"errors"
	"gorm.io/gorm"
	"time"
)

var EmployeeService = newEmployeeService()

func newEmployeeService() *employeeService {
	return &employeeService{}
}

type employeeService struct {
}

func (cc *employeeService) GetEmployeeByToken(token string) (*model.Employee, int) {
	claims, _ := util.ParseToken(token)
	Employee := new(model.Employee)
	err := global.DB.Where("account = ?", claims.Account).Take(Employee).Error
	if err != nil {
		return nil, e.ERROR_AUTH
	}
	return Employee, e.SUCCESS
}
func (cc *employeeService) RegisterEmployee(req *employee.RegisterRequest, token string) int {
	if req.Account == "" || req.Password == "" {
		return e.INVALID_PARAMS
	}
	RegisterUser, errcode := cc.GetEmployeeByToken(token)
	if errcode != e.SUCCESS {
		return errcode
	}
	user := &model.Employee{
		Account:    req.Account,
		Password:   req.Password,
		Phone:      "18877776666",
		CreateUser: RegisterUser.Id,
		UpdateUser: RegisterUser.Id,
		CreateTime: jsontime.JSONTime{Time: time.Now()},
		UpdateTime: jsontime.JSONTime{Time: time.Now()},
	}
	err := global.DB.Create(user).Error
	if err != nil {
		return e.ERROR_DATABASE
	}
	return e.SUCCESS
}

func (cc *employeeService) LoginEmployee(req *employee.LoginRequest) (*employee.LoginResponse, int) {
	if req.Account == "" || req.Password == "" {
		return nil, e.INVALID_PARAMS
	}
	user := new(model.Employee)
	err := global.DB.Where("account = ?", req.Account).Take(user).Error
	if err != nil {
		return nil, e.ERROR_DATABASE
	}
	if req.Password != user.Password {
		return nil, e.ERROR_PASSWORD
	}
	token, err := util.GenerateToken(req.Account, req.Password)
	if err != nil {
		return nil, e.ERROR_AUTH_TOKEN
	}
	return &employee.LoginResponse{
		Id:      user.Id,
		Token:   token,
		Account: user.Account,
	}, e.SUCCESS
}

func (cc *employeeService) EditEmployeePassword(req *employee.EditPasswordRequest, token string) int {
	if req.OldPwd == "" || req.NewPwd == "" {
		return e.INVALID_PARAMS
	}
	Employee, errcode := cc.GetEmployeeByToken(token)
	if errcode != e.SUCCESS {
		return errcode
	}
	if Employee.Password != req.OldPwd {
		return e.ERROR_PASSWORD
	}
	err := global.DB.Model(Employee).Updates(map[string]interface{}{
		"password":    req.NewPwd,
		"update_user": Employee.Id,
	}).Error
	if err != nil {
		return e.ERROR_DATABASE
	}
	return e.SUCCESS
}

func (cc *employeeService) AddEmployee(req *employee.AddEditRequest, token string) int {
	// 参数检验部分省略
	Employee, errcode := cc.GetEmployeeByToken(token)
	if errcode != e.SUCCESS {
		return errcode
	}
	user := &model.Employee{
		Name:       req.Name,
		Account:    req.Account,
		Password:   req.Password,
		Phone:      req.Phone,
		Age:        req.Age,
		Gender:     req.Gender,
		Pic:        req.Pic,
		CreateUser: Employee.Id,
		UpdateUser: Employee.Id,
		Status:     1,
		CreateTime: jsontime.JSONTime{Time: time.Now()},
		UpdateTime: jsontime.JSONTime{Time: time.Now()},
	}
	err := global.DB.Create(user).Error
	if err != nil {
		return e.ERROR_ALREADY_EXIST
	}
	return e.SUCCESS
}

func (cc *employeeService) EditEmployee(req *employee.AddEditRequest, token string) int {
	// 参数检验部分省略
	Employee, errcode := cc.GetEmployeeByToken(token)
	if errcode != e.SUCCESS {
		return errcode
	}
	user := &model.Employee{
		Id:         req.Id,
		Name:       req.Name,
		Account:    req.Account,
		Password:   req.Password,
		Phone:      req.Phone,
		Age:        req.Age,
		Gender:     req.Gender,
		Pic:        req.Pic,
		UpdateUser: Employee.Id,
	}
	err := global.DB.Updates(user).Error
	if err != nil {
		return e.ERROR_ALREADY_EXIST
	}
	return e.SUCCESS
}

func (cc *employeeService) GetEmployeeById(id int, token string) (*model.Employee, int) {
	Employee, errcode := cc.GetEmployeeByToken(token)
	if errcode != e.SUCCESS {
		return nil, errcode
	}
	if Employee.Id == id {
		return Employee, e.SUCCESS
	}
	if Employee.Account == "admin" {
		Employee := new(model.Employee)
		err := global.DB.Where("id = ?", id).First(Employee).Error
		if err != nil {
			return nil, e.ERROR_DATABASE
		}
		return Employee, e.SUCCESS
	} else {
		return nil, e.ERROR_ACCESS
	}
}

func (cc *employeeService) GetEmployeeList(req *employee.PageRequest) (*employee.PageResponse, int) {
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)
	var Employees []*model.Employee
	if req.Name != "" {
		err := global.DB.Where("name = ?", req.Name).Offset(offset).Limit(limit).Find(&Employees).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, e.ERROR_DATABASE
		}
	} else {
		err := global.DB.Offset(offset).Limit(limit).Find(&Employees).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, e.ERROR_DATABASE
		}
	}
	return &employee.PageResponse{
		Total:   len(Employees),
		Records: Employees,
	}, e.SUCCESS
}

func (cc *employeeService) EditEmployeeStatus(id int, token string) int {
	admin, errcode := cc.GetEmployeeByToken(token)
	if errcode != e.SUCCESS {
		return e.ERROR_ACCESS
	}
	Employee := new(model.Employee)
	err := global.DB.Where("id = ?", id).First(Employee).Error
	if err != nil {
		return e.ERROR_DATABASE
	}
	var status int
	if Employee.Status == 0 {
		status = 1
	}
	err = global.DB.Model(Employee).Updates(map[string]interface{}{
		"status":      status,
		"update_user": admin.Id,
	}).Error
	if err != nil {
		return e.ERROR_DATABASE
	}
	return e.SUCCESS
}

func (cc *employeeService) DeleteEmployee(id int, token string) int {
	admin, errcode := cc.GetEmployeeByToken(token)
	if errcode != e.SUCCESS {
		return e.ERROR_ACCESS
	}
	if admin.Account != "admin" {
		return e.ERROR_ACCESS
	}
	Employee := new(model.Employee)
	err := global.DB.Where("id = ?", id).First(Employee).Error
	if err != nil {
		return e.ERROR_DATABASE
	}
	err = global.DB.Delete(Employee).Error
	if err != nil {
		return e.ERROR_DATABASE
	}
	return e.SUCCESS
}
