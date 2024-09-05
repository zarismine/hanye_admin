package api

import (
	"admin_app/pkg/e"
	"admin_app/service"
	"admin_app/web"
	"github.com/kataras/iris/v12"
)

type WorkspaceController struct {
	Ctx iris.Context
}

func (cc *WorkspaceController) GetBy(target string) *web.JsonResult {
	if target == "overviewDishes" {
		data, errcode := service.WorkspaceService.OverviewDishes()
		if errcode != e.SUCCESS {
			return web.JsonError(errcode)
		}
		return web.JsonData(data)
	}
	if target == "overviewSetmeals" {
		data, errcode := service.WorkspaceService.OverviewDishes()
		if errcode != e.SUCCESS {
			return web.JsonError(errcode)
		}
		return web.JsonData(data)
	}
	if target == "overviewOrders" {
		data, errcode := service.WorkspaceService.OverviewOrders()
		if errcode != e.SUCCESS {
			return web.JsonError(errcode)
		}
		return web.JsonData(data)
	}
	if target == "businessData" {
		data, errcode := service.WorkspaceService.BusinessData()
		if errcode != e.SUCCESS {
			return web.JsonError(errcode)
		}
		return web.JsonData(data)
	}
	return web.JsonSuccess()
}
