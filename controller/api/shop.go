package api

import (
	"admin_app/global"
	"admin_app/pkg/e"
	"admin_app/web"
	"github.com/kataras/iris/v12"
)

type ShopController struct {
	Ctx iris.Context
}

func (cc *ShopController) GetStatus() *web.JsonResult {
	return web.JsonData(global.ShopStatus)
}

func (cc *ShopController) PutBy(status int) *web.JsonResult {
	if status != 0 && status != 1 {
		return web.JsonError(e.ERROR)
	}
	global.ShopStatus = status
	return web.JsonSuccess()
}
