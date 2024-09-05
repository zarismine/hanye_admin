package api

import (
	"admin_app/pkg/e"
	"admin_app/service"
	"admin_app/types/setmeal"
	"admin_app/web"
	"github.com/kataras/iris/v12"
)

type SetmealController struct {
	Ctx iris.Context
}

func (cc *SetmealController) Post() *web.JsonResult {
	req := new(setmeal.AddRequest)
	cc.Ctx.ReadJSON(req)
	token := cc.Ctx.GetHeader("Authorization")
	errcode := service.SetmealService.AddSetmeal(req, token)
	if errcode != e.SUCCESS {
		return web.JsonError(errcode)
	}
	return web.JsonSuccess()
}

func (cc *SetmealController) GetList() *web.JsonResult {
	req := new(setmeal.CategorySearchRequest)
	cc.Ctx.ReadQuery(req)
	data, errcode := service.SetmealService.CategorySearchSetmealById(req.CategoryId)
	if errcode != e.SUCCESS {
		return web.JsonError(errcode)
	}
	return web.JsonData(data)
}

func (cc *SetmealController) GetPage() *web.JsonResult {
	req := new(setmeal.PageRequest)
	req.Status = -1
	cc.Ctx.ReadQuery(req)
	data, errcode := service.SetmealService.GetSetmealPage(req)
	if errcode != e.SUCCESS {
		return web.JsonError(errcode)
	}
	return web.JsonData(data)
}

func (cc *SetmealController) GetBy(id int) *web.JsonResult {
	data, errcode := service.SetmealService.GetSetmealById(id)
	if errcode != e.SUCCESS {
		return web.JsonError(errcode)
	}
	return web.JsonData(data)
}

func (cc *SetmealController) PutStatusBy(id int) *web.JsonResult {
	token := cc.Ctx.GetHeader("Authorization")
	errcode := service.SetmealService.EditSetmealStatusById(id, token)
	if errcode != e.SUCCESS {
		return web.JsonError(errcode)
	}
	return web.JsonSuccess()
}

func (cc *SetmealController) Put() *web.JsonResult {
	req := new(setmeal.EditRequest)
	cc.Ctx.ReadJSON(req)
	token := cc.Ctx.GetHeader("Authorization")
	errcode := service.SetmealService.EditSetmeal(req, token)
	if errcode != e.SUCCESS {
		return web.JsonError(errcode)
	}
	return web.JsonSuccess()
}

func (cc *SetmealController) Delete() *web.JsonResult {
	req := new(setmeal.DeleteRequest)
	cc.Ctx.ReadQuery(req)
	errcode := service.SetmealService.DeleteSetmealByIds(req)
	if errcode != e.SUCCESS {
		return web.JsonError(errcode)
	}
	return web.JsonSuccess()
}
