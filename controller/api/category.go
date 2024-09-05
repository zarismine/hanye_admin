package api

import (
	"admin_app/pkg/e"
	"admin_app/service"
	"admin_app/types/category"
	"admin_app/web"
	"github.com/kataras/iris/v12"
)

type CategoryController struct {
	Ctx iris.Context
}

func (cc *CategoryController) Post() *web.JsonResult {
	req := new(category.AddRequest)
	cc.Ctx.ReadJSON(req)
	token := cc.Ctx.GetHeader("Authorization")
	errcode := service.CategoryService.AddCategory(req, token)
	if errcode != e.SUCCESS {
		return web.JsonError(errcode)
	}
	return web.JsonSuccess()
}

func (cc *CategoryController) GetPage() *web.JsonResult {
	req := new(category.PageRequest)
	cc.Ctx.ReadQuery(req)
	data, errcode := service.CategoryService.PageCategory(req)
	if errcode != e.SUCCESS {
		return web.JsonError(errcode)
	}
	return web.JsonData(data)
}

func (cc *CategoryController) GetBy(id int) *web.JsonResult {
	data, errcode := service.CategoryService.GetDetailById(id)
	if errcode != e.SUCCESS {
		return web.JsonError(errcode)
	}
	return web.JsonData(data)
}

func (cc *CategoryController) PutStatusBy(id int) *web.JsonResult {
	token := cc.Ctx.GetHeader("Authorization")
	errcode := service.CategoryService.EditCategoryStatus(id, token)
	if errcode != e.SUCCESS {
		return web.JsonError(errcode)
	}
	return web.JsonSuccess()
}

func (cc *CategoryController) Put() *web.JsonResult {
	req := new(category.EditRequest)
	cc.Ctx.ReadJSON(req)
	token := cc.Ctx.GetHeader("Authorization")
	errcode := service.CategoryService.EditCategory(req, token)
	if errcode != e.SUCCESS {
		return web.JsonError(errcode)
	}
	return web.JsonSuccess()
}

func (cc *CategoryController) DeleteBy(id int) *web.JsonResult {
	errcode := service.CategoryService.DeleteCategoryById(id)
	if errcode != e.SUCCESS {
		return web.JsonError(errcode)
	}
	return web.JsonSuccess()
}

func (cc *CategoryController) GetList() *web.JsonResult {
	req := new(category.GetListRequest)
	cc.Ctx.ReadQuery(req)
	data, errcode := service.CategoryService.GetListByType(req)
	if errcode != e.SUCCESS {
		return web.JsonError(errcode)
	}
	return web.JsonData(data)
}
