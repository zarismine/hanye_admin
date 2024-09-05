package router

import (
	"admin_app/controller/api"
	"admin_app/controller/user"
	"admin_app/middleware/jwt"
	"github.com/kataras/iris/v12"

	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/mvc"
)

func crs(ctx iris.Context) {
	origin := ctx.GetHeader("Origin")
	ctx.Header("Access-Control-Allow-Origin", origin)
	ctx.Header("Access-Control-Allow-Credentials", "true")
	if ctx.Method() == iris.MethodOptions {
		ctx.Header("Access-Control-Allow-Methods", "POST, PUT, PATCH, DELETE")
		ctx.Header("Access-Control-Allow-Headers", "Access-Control-Allow-Origin,Content-Type,X-Requested-With,Token")
		ctx.Header("Access-Control-Max-Age", "86400")
		ctx.StatusCode(iris.StatusNoContent)
		return
	}
	ctx.Next()
}

func NewServer() {
	app := iris.New()
	app.Use(recover.New())
	app.Use(logger.New())
	app.UseRouter(crs)
	mvc.Configure(app.Party("/admin", jwt.JWT), func(m *mvc.Application) {
		m.Party("/employee").Handle(new(api.EmployeeController))
		m.Party("/category").Handle(new(api.CategoryController))
		m.Party("/setmeal").Handle(new(api.SetmealController))
		m.Party("/dish").Handle(new(api.DishController))
		m.Party("/shop").Handle(new(api.ShopController))
		m.Party("/workspace").Handle(new(api.WorkspaceController))
		m.Party("/order").Handle(new(api.OrderController))
	})
	mvc.Configure(app.Party("/user"), func(m *mvc.Application) {
		m.Party("/shop").Handle(new(user.ShopController))
	})
	app.Run(iris.Addr(":19999"))
}
