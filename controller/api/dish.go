package api

import (
	"admin_app/pkg/e"
	"admin_app/service"
	"admin_app/types/dish"
	"admin_app/web"
	"github.com/kataras/iris/v12"
)

type DishController struct {
	Ctx iris.Context
}

func (cc *DishController) Post() *web.JsonResult {
	req := new(dish.AddRequest)
	cc.Ctx.ReadJSON(req)
	token := cc.Ctx.GetHeader("Authorization")
	errcode := service.DishService.AddDish(req, token)
	if errcode != e.SUCCESS {
		return web.JsonError(errcode)
	}
	return web.JsonSuccess()
}

func (cc *DishController) GetPage() *web.JsonResult {
	req := new(dish.PageRequest)
	req.Status = -1
	cc.Ctx.ReadQuery(req)
	data, errcode := service.DishService.GetDishPage(req)
	if errcode != e.SUCCESS {
		return web.JsonError(errcode)
	}
	return web.JsonData(data)
}

func (cc *DishController) GetBy(id int) *web.JsonResult {
	data, errcode := service.DishService.GetDishAndFlavorBy(id)
	if errcode != e.SUCCESS {
		return web.JsonError(errcode)
	}
	return web.JsonData(data)
}

func (cc *DishController) PutStatusBy(id int) *web.JsonResult {
	token := cc.Ctx.GetHeader("Authorization")
	errcode := service.DishService.EditDishStatusById(id, token)
	if errcode != e.SUCCESS {
		return web.JsonError(errcode)
	}
	return web.JsonSuccess()
}

func (cc *DishController) Put() *web.JsonResult {
	req := new(dish.EditRequest)
	cc.Ctx.ReadJSON(req)
	token := cc.Ctx.GetHeader("Authorization")
	errcode := service.DishService.EditDish(req, token)
	if errcode != e.SUCCESS {
		return web.JsonError(errcode)
	}
	return web.JsonSuccess()
}

func (cc *DishController) Delete() *web.JsonResult {
	req := new(dish.DeleteRequest)
	cc.Ctx.ReadQuery(req)
	errcode := service.DishService.DeleteDishByIds(req)
	if errcode != e.SUCCESS {
		return web.JsonError(errcode)
	}
	return web.JsonSuccess()
}
