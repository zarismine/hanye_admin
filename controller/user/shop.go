package user

import (
	"admin_app/global"
	"admin_app/web"
	"github.com/kataras/iris/v12"
)

type ShopController struct {
	Ctx iris.Context
}

func (cc *ShopController) GetStatus() *web.JsonResult {
	return web.JsonData(global.ShopStatus)
}
