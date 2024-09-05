package api

import (
	"admin_app/pkg/e"
	"admin_app/service"
	"admin_app/types/order"
	"admin_app/web"
	"github.com/kataras/iris/v12"
)

type OrderController struct {
	Ctx iris.Context
}

func (m *OrderController) GetPage() *web.JsonResult {
	req := new(order.PageReq)
	_ = m.Ctx.ReadQuery(req)
	data, errcode := service.OrderService.PageList(req)
	if errcode != e.SUCCESS {
		return web.JsonError(errcode)
	}
	return web.JsonData(data)
}

func (m *OrderController) GetStatistics() *web.JsonResult {
	data, errcode := service.OrderService.Statistics()
	if errcode != e.SUCCESS {
		return web.JsonError(errcode)
	}
	return web.JsonData(data)
}

func (m *OrderController) PutConfirm() *web.JsonResult {
	req := new(order.ConfirmReq)
	_ = m.Ctx.ReadJSON(req)
	errcode := service.OrderService.Confirm(req)
	if errcode != e.SUCCESS {
		return web.JsonError(errcode)
	}
	return web.JsonSuccess()
}

func (m *OrderController) PutReject() *web.JsonResult {
	req := new(order.RejectReq)
	_ = m.Ctx.ReadJSON(req)
	errcode := service.OrderService.Reject(req)
	if errcode != e.SUCCESS {
		return web.JsonError(errcode)
	}
	return web.JsonSuccess()
}

func (m *OrderController) PutDeliveryBy(id int) *web.JsonResult {
	errcode := service.OrderService.Delivery(id)
	if errcode != e.SUCCESS {
		return web.JsonError(errcode)
	}
	return web.JsonSuccess()
}

func (m *OrderController) PutCompleteBy(id int) *web.JsonResult {
	errcode := service.OrderService.Complete(id)
	if errcode != e.SUCCESS {
		return web.JsonError(errcode)
	}
	return web.JsonSuccess()
}

func (m *OrderController) GetDetailsBy(id int) *web.JsonResult {
	data, errcode := service.OrderService.Details(id)
	if errcode != e.SUCCESS {
		return web.JsonError(errcode)
	}
	return web.JsonData(data)
}
