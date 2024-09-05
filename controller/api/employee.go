package api

import (
	"admin_app/pkg/e"
	"admin_app/service"
	"admin_app/types/employee"
	"admin_app/web"
	"github.com/kataras/iris/v12"
)

type EmployeeController struct {
	Ctx iris.Context
}

func (m *EmployeeController) PostRegister() *web.JsonResult {
	req := new(employee.RegisterRequest)
	m.Ctx.ReadJSON(req)
	token := m.Ctx.GetHeader("Authorization")
	errcode := service.EmployeeService.RegisterEmployee(req, token)
	if errcode != e.SUCCESS {
		return web.JsonError(errcode)
	}
	return web.JsonSuccess()
}

func (m *EmployeeController) PostLogin() *web.JsonResult {
	req := new(employee.LoginRequest)
	m.Ctx.ReadJSON(req)
	data, errcode := service.EmployeeService.LoginEmployee(req)
	if errcode != e.SUCCESS {
		return web.JsonError(errcode)
	}
	return web.JsonData(data)
}

func (m *EmployeeController) PutFixpwd() *web.JsonResult {
	req := new(employee.EditPasswordRequest)
	m.Ctx.ReadJSON(req)
	token := m.Ctx.GetHeader("Authorization")
	errcode := service.EmployeeService.EditEmployeePassword(req, token)
	if errcode != e.SUCCESS {
		return web.JsonError(errcode)
	}
	return web.JsonSuccess()
}

func (m *EmployeeController) PostAdd() *web.JsonResult {
	req := new(employee.AddEditRequest)
	m.Ctx.ReadJSON(req)
	token := m.Ctx.GetHeader("Authorization")
	errcode := service.EmployeeService.AddEmployee(req, token)
	if errcode != e.SUCCESS {
		return web.JsonError(errcode)
	}
	return web.JsonSuccess()
}

func (m *EmployeeController) PutUpdate() *web.JsonResult {
	req := new(employee.AddEditRequest)
	m.Ctx.ReadJSON(req)
	token := m.Ctx.GetHeader("Authorization")
	errcode := service.EmployeeService.EditEmployee(req, token)
	if errcode != e.SUCCESS {
		return web.JsonError(errcode)
	}
	return web.JsonSuccess()
}

func (m *EmployeeController) GetBy(id int) *web.JsonResult {
	token := m.Ctx.GetHeader("Authorization")
	data, errcode := service.EmployeeService.GetEmployeeById(id, token)
	if errcode != e.SUCCESS {
		return web.JsonError(errcode)
	}
	return web.JsonData(data)
}

func (m *EmployeeController) GetPage() *web.JsonResult {
	req := new(employee.PageRequest)
	m.Ctx.ReadQuery(req)
	data, errcode := service.EmployeeService.GetEmployeeList(req)
	if errcode != e.SUCCESS {
		return web.JsonError(errcode)
	}
	return web.JsonData(data)
}

func (m *EmployeeController) PutStatusBy(id int) *web.JsonResult {
	token := m.Ctx.GetHeader("Authorization")
	errcode := service.EmployeeService.EditEmployeeStatus(id, token)
	if errcode != e.SUCCESS {
		return web.JsonError(errcode)
	}
	return web.JsonSuccess()
}

func (m *EmployeeController) DeleteDeleteBy(id int) *web.JsonResult {
	token := m.Ctx.GetHeader("Authorization")
	errcode := service.EmployeeService.DeleteEmployee(id, token)
	if errcode != e.SUCCESS {
		return web.JsonError(errcode)
	}
	return web.JsonSuccess()
}
