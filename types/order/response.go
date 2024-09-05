package order

import "admin_app/model"

type NewOrders struct {
	model.Orders
	Dishes string `json:"orderDishes"`
}

type PageResp struct {
	Total   int64        `json:"total"`
	Records []*NewOrders `json:"records"`
}

type Statistics struct {
	ToBeConfirmed      int64 `json:"toBeConfirmed"`
	Confirmed          int64 `json:"confirmed"`
	DeliveryInProgress int64 `json:"deliveryInProgress"`
}

type ConfirmResp struct {
}

type DetailResp struct {
	NewOrders
	Details []*model.OrderDetail `json:"orderDetailList"`
}
