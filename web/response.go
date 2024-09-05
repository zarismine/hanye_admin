package web

import "admin_app/pkg/e"

type JsonResult struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func JsonData(data interface{}) *JsonResult {
	return &JsonResult{
		Code: e.SUCCESS,
		Data: data,
	}
}

func JsonSuccess() *JsonResult {
	return &JsonResult{
		Code: e.SUCCESS,
		Data: nil,
	}
}

func JsonError(errcode int) *JsonResult {
	if errcode == e.SUCCESS {
		return JsonSuccess()
	}

	return &JsonResult{
		Code:    errcode,
		Message: e.GetMsg(errcode),
		Data:    nil,
	}
}
