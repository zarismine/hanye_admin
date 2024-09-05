package jwt

import (
	"admin_app/pkg/e"
	util "admin_app/pkg/utils"
	"admin_app/web"
	"github.com/kataras/iris/v12"
	"time"
)

type PassPath struct {
	Path   string
	Method string
}

var Pass = []PassPath{
	{Path: "/admin/employee/login", Method: "POST"},
	{Path: "/admin/employee/register", Method: "POST"},
}

func containsPathMethod(pass []PassPath, path, method string) bool {
	for _, pp := range pass {
		if pp.Path == path && pp.Method == method {
			return true
		}
	}
	return false
}

func JWT(ctx iris.Context) {
	var code = e.SUCCESS
	p := ctx.Path()
	m := ctx.Method()
	if containsPathMethod(Pass, p, m) {
		ctx.Next()
		return
	}
	token := ctx.GetHeader("Authorization")
	if token == "" {
		code = e.INVALID_PARAMS
	} else {
		claims, err := util.ParseToken(token)
		if err != nil {
			code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
		} else if time.Now().Unix() > claims.ExpiresAt {
			code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
		}
	}
	if code != e.SUCCESS {
		ctx.JSON(web.JsonError(code))
		return
	}
	ctx.Next()
}
