package service

import (
	"admin_app/global"
	"admin_app/model"
	"admin_app/model/constant"
	"admin_app/pkg/e"
	"admin_app/types/workspace"
	"time"
)

var WorkspaceService = newWorkspaceService()

func newWorkspaceService() *workspaceService {
	return &workspaceService{}
}

type workspaceService struct {
}

func (cc *workspaceService) OverviewDishes() (*workspace.Response, int) {
	var CountUp, CountDown int64
	var err error
	err = global.DB.Table("dish").Where("status = ?", constant.StatusUp).Count(&CountUp).Error
	if err != nil {
		return nil, e.ERROR_DATABASE
	}
	err = global.DB.Table("dish").Where("status = ?", constant.StatusDown).Count(&CountDown).Error
	if err != nil {
		return nil, e.ERROR_DATABASE
	}
	return &workspace.Response{
		CountUp:   CountUp,
		CountDown: CountDown,
	}, e.SUCCESS
}

func (cc *workspaceService) OverviewSetmeals() (*workspace.Response, int) {
	var CountUp, CountDown int64
	var err error
	err = global.DB.Table("setmeal").Where("status = ?", constant.StatusUp).Count(&CountUp).Error
	if err != nil {
		return nil, e.ERROR_DATABASE
	}
	err = global.DB.Table("setmeal").Where("status = ?", constant.StatusDown).Count(&CountDown).Error
	if err != nil {
		return nil, e.ERROR_DATABASE
	}
	return &workspace.Response{
		CountUp:   CountUp,
		CountDown: CountDown,
	}, e.SUCCESS
}

func (cc *workspaceService) OverviewOrders() (*workspace.OrderResponse, int) {
	var waitingOrders, deliveredOrders, completedOrders, cancelledOrders, allOrders int64
	var err error
	now := time.Now()
	todayMidnight := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	err = global.DB.Table("orders").Where("status = ? and order_time >= ?", constant.StatusWaitingOrders, todayMidnight).Count(&waitingOrders).Error
	if err != nil {
		return nil, e.ERROR_DATABASE
	}
	err = global.DB.Table("orders").Where("status = ? and order_time >= ?", constant.StatusDeliveredOrders, todayMidnight).Count(&deliveredOrders).Error
	if err != nil {
		return nil, e.ERROR_DATABASE
	}
	err = global.DB.Table("orders").Where("status = ? and order_time >= ?", constant.StatusCompletedOrders, todayMidnight).Count(&completedOrders).Error
	if err != nil {
		return nil, e.ERROR_DATABASE
	}
	err = global.DB.Table("orders").Where("status = ? and order_time >= ?", constant.StatusCancelledOrders, todayMidnight).Count(&cancelledOrders).Error
	if err != nil {
		return nil, e.ERROR_DATABASE
	}
	err = global.DB.Table("orders").Where("order_time >= ?", todayMidnight).Count(&allOrders).Error
	if err != nil {
		return nil, e.ERROR_DATABASE
	}
	return &workspace.OrderResponse{
		WaitingOrders:   waitingOrders,
		DeliveredOrders: deliveredOrders,
		CompletedOrders: completedOrders,
		CancelledOrders: cancelledOrders,
		AllOrders:       allOrders,
	}, e.SUCCESS
}

func (cc *workspaceService) BusinessData() (*workspace.BusinessResponse, int) {
	var (
		allOrders           int64
		turnover            float64
		validOrderCount     int64
		orderCompletionRate float64
		unitPrice           float64
		newUsers            int64
	)
	now := time.Now()
	todayMidnight := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	DB := global.DB.Table("orders").Where("order_time >= ?", todayMidnight)
	var err error
	err = DB.Count(&allOrders).Error
	if err != nil {
		return nil, e.ERROR_DATABASE
	}
	err = DB.Where("status = ?", constant.StatusCompletedOrders).Count(&validOrderCount).Error
	if err != nil {
		return nil, e.ERROR_DATABASE
	}
	var newOrders []*model.Orders
	err = DB.Find(&newOrders).Error
	if err != nil {
		return nil, e.ERROR_DATABASE
	}
	for _, v := range newOrders {
		turnover += v.Amount
	}
	if allOrders == 0 {
		orderCompletionRate = 0
	} else {
		orderCompletionRate = float64(validOrderCount) / float64(allOrders)
	}
	if validOrderCount == 0 {
		unitPrice = 0
	} else {
		unitPrice = turnover / float64(validOrderCount)
	}
	err = global.DB.Table("user").Where("create_time >= ?", todayMidnight).Count(&newUsers).Error
	if err != nil {
		return nil, e.ERROR_DATABASE
	}
	return &workspace.BusinessResponse{
		Turnover:            turnover,
		ValidOrderCount:     validOrderCount,
		OrderCompletionRate: orderCompletionRate,
		UnitPrice:           unitPrice,
		NewUsers:            newUsers,
	}, e.SUCCESS
}
