package service

import (
	"admin_app/global"
	"admin_app/model"
	"admin_app/model/constant"
	"admin_app/model/jsontime"
	"admin_app/pkg/e"
	"admin_app/types/order"
	"github.com/jinzhu/copier"

	"errors"
	"gorm.io/gorm"
	"time"
)

var OrderService = newOrderService()

func newOrderService() *orderService {
	return &orderService{}
}

type orderService struct {
}

func (cc *orderService) PageList(req *order.PageReq) (*order.PageResp, int) {
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)
	var orders []*model.Orders
	DB := global.DB.Table("orders")
	if req.BeginTime != "" {
		beginTime, _ := time.Parse(time.DateTime, req.BeginTime)
		DB = DB.Where("order_time >= ?", beginTime)
	}
	if req.EndTime != "" {
		endTime, _ := time.Parse(time.DateTime, req.EndTime)
		DB = DB.Where("order_time < ?", endTime)
	}
	if req.Phone != "" {
		DB = DB.Where("phone = ?", req.Phone)
	}
	if req.Number != "" {
		DB = DB.Where("number = ?", req.Number)
	}
	if req.Status != 0 {
		DB = DB.Where("status = ?", req.Status)

	}
	//if req.UserID != "" {
	//	err := DB.Where("number = ?", req.Number).Offset(offset).Limit(limit).Find(&orders).Error
	//	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
	//		return nil, e.ERROR_DATABASE
	//	}
	//}
	var total int64
	DB.Count(&total)
	err := DB.Offset(offset).Limit(limit).Order("order_time desc").Find(&orders).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, e.SUCCESS
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, e.ERROR_DATABASE
	}
	var newOrders []*order.NewOrders
	_ = copier.Copy(&newOrders, orders)
	for _, v := range newOrders {
		orderDetail := new([]*model.OrderDetail)
		err = global.DB.Where("order_id = ?", v.Id).Find(orderDetail).Error
		if err != nil {
			return nil, e.ERROR_DATABASE
		}
		for _, vv := range *orderDetail {
			v.Dishes += vv.Name + vv.DishFlavor + "\n"
		}
	}
	return &order.PageResp{
		Total:   total,
		Records: newOrders,
	}, e.SUCCESS
}

func (cc *orderService) Statistics() (*order.Statistics, int) {
	var toBeConfirmed, confirmed, deliveryInProgress int64
	var err error
	err = global.DB.Table("orders").Where("status = ?", constant.StatusWaitingOrders).Count(&toBeConfirmed).Error
	if err != nil {
		return nil, e.ERROR_DATABASE
	}
	err = global.DB.Table("orders").Where("status = ?", constant.StatusDeliveredOrders).Count(&confirmed).Error
	if err != nil {
		return nil, e.ERROR_DATABASE
	}
	err = global.DB.Table("orders").Where("status = ?", constant.StatusDeliveringOrders).Count(&deliveryInProgress).Error
	if err != nil {
		return nil, e.ERROR_DATABASE
	}

	return &order.Statistics{
		ToBeConfirmed:      toBeConfirmed,
		Confirmed:          confirmed,
		DeliveryInProgress: deliveryInProgress,
	}, e.SUCCESS
}

func (cc *orderService) Confirm(req *order.ConfirmReq) int {
	var err error
	var orderTarget model.Orders
	err = global.DB.Table("orders").Where("id = ?", req.Id).First(&orderTarget).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return e.ERROR_EXIST
	}
	if err != nil {
		return e.ERROR_DATABASE
	}
	err = global.DB.Model(&orderTarget).UpdateColumn("status", constant.StatusDeliveredOrders).Error
	if err != nil {
		return e.ERROR_DATABASE
	}
	return e.SUCCESS
}

func (cc *orderService) Reject(req *order.RejectReq) int {
	var err error
	var orderTarget model.Orders
	err = global.DB.Table("orders").Where("id = ?", req.Id).First(&orderTarget).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return e.ERROR_EXIST
	}
	if err != nil {
		return e.ERROR_DATABASE
	}
	err = global.DB.Model(&orderTarget).UpdateColumns(map[string]interface{}{
		"status":           constant.StatusCancelledOrders,
		"rejection_reason": req.RejectionReason,
		"cancel_time":      jsontime.JSONTime{Time: time.Now()},
	}).Error
	if err != nil {
		return e.ERROR_DATABASE
	}
	return e.SUCCESS
}

func (cc *orderService) Delivery(id int) int {
	var err error
	var orderTarget model.Orders
	err = global.DB.Table("orders").Where("id = ?", id).First(&orderTarget).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return e.ERROR_EXIST
	}
	if err != nil {
		return e.ERROR_DATABASE
	}
	err = global.DB.Model(&orderTarget).UpdateColumn("status", constant.StatusDeliveringOrders).Error
	if err != nil {
		return e.ERROR_DATABASE
	}
	return e.SUCCESS
}

func (cc *orderService) Complete(id int) int {
	var err error
	var orderTarget model.Orders
	err = global.DB.Table("orders").Where("id = ?", id).First(&orderTarget).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return e.ERROR_EXIST
	}
	if err != nil {
		return e.ERROR_DATABASE
	}
	err = global.DB.Model(&orderTarget).UpdateColumns(map[string]interface{}{
		"status":        constant.StatusCompletedOrders,
		"delivery_time": jsontime.JSONTime{Time: time.Now()},
	}).Error
	if err != nil {
		return e.ERROR_DATABASE
	}
	return e.SUCCESS
}

func (cc *orderService) Details(id int) (*order.DetailResp, int) {
	var orderTarget model.Orders
	err := global.DB.Table("orders").Where("id = ?", id).First(&orderTarget).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, e.SUCCESS
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, e.ERROR_DATABASE
	}
	var newOrders order.NewOrders
	_ = copier.Copy(&newOrders, orderTarget)
	orderDetail := new([]*model.OrderDetail)
	err = global.DB.Where("order_id = ?", newOrders.Id).Find(orderDetail).Error
	if err != nil {
		return nil, e.ERROR_DATABASE
	}
	for _, vv := range *orderDetail {
		newOrders.Dishes += vv.Name + vv.DishFlavor + "\n"
	}
	return &order.DetailResp{
		NewOrders: newOrders,
		Details:   *orderDetail,
	}, e.SUCCESS
}
