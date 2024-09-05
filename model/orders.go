package model

import (
	"admin_app/model/jsontime"
)

type Orders struct {
	Id                    int64             `gorm:"column:id;type:bigint;comment:主键;primaryKey;" json:"id"`                                                             // 主键
	Number                string            `gorm:"column:number;type:varchar(50);comment:订单号;" json:"number"`                                                          // 订单号
	Status                int32             `gorm:"column:status;type:int;comment:订单状态 1待付款 2待接单 3已接单 4派送中 5已完成 6已取消 7退款;not null;default:1;" json:"status"`            // 订单状态 1待付款 2待接单 3已接单 4派送中 5已完成 6已取消 7退款
	UserId                int64             `gorm:"column:user_id;type:bigint;comment:下单用户;not null;" json:"userId"`                                                    // 下单用户
	AddressBookId         int64             `gorm:"column:address_book_id;type:bigint;comment:地址id;not null;" json:"addressBookId"`                                     // 地址id
	OrderTime             jsontime.JSONTime `gorm:"column:order_time;type:datetime;comment:下单时间;not null;" json:"orderTime"`                                            // 下单时间
	CheckoutTime          jsontime.JSONTime `gorm:"column:checkout_time;type:datetime;comment:结账时间;" json:"checkoutTime"`                                               // 结账时间
	PayMethod             int32             `gorm:"column:pay_method;type:int;comment:支付方式 1微信,2支付宝;not null;default:1;" json:"payMethod"`                              // 支付方式 1微信,2支付宝
	PayStatus             int32             `gorm:"column:pay_status;type:tinyint;comment:支付状态 0未支付 1已支付 2退款;not null;default:0;" json:"payStatus"`                     // 支付状态 0未支付 1已支付 2退款
	Amount                float64           `gorm:"column:amount;type:decimal(10, 2);comment:实收金额;not null;" json:"amount"`                                             // 实收金额
	Remark                string            `gorm:"column:remark;type:varchar(100);comment:备注;" json:"remark"`                                                          // 备注
	Phone                 string            `gorm:"column:phone;type:varchar(11);comment:手机号;" json:"phone"`                                                            // 手机号
	Address               string            `gorm:"column:address;type:varchar(255);comment:地址;" json:"address"`                                                        // 地址
	UserName              string            `gorm:"column:user_name;type:varchar(32);comment:用户名称;" json:"userName"`                                                    // 用户名称
	Consignee             string            `gorm:"column:consignee;type:varchar(32);comment:收货人;" json:"consignee"`                                                    // 收货人
	CancelReason          string            `gorm:"column:cancel_reason;type:varchar(255);comment:订单取消原因;" json:"cancelReason"`                                         // 订单取消原因
	RejectionReason       string            `gorm:"column:rejection_reason;type:varchar(255);comment:订单拒绝原因;" json:"rejectionReason"`                                   // 订单拒绝原因
	CancelTime            jsontime.JSONTime `gorm:"column:cancel_time;type:datetime;comment:订单取消时间;" json:"cancelTime"`                                                 // 订单取消时间
	EstimatedDeliveryTime jsontime.JSONTime `gorm:"column:estimated_delivery_time;type:datetime;comment:预计送达时间;" json:"estimatedDeliveryTime"`                          // 预计送达时间
	DeliveryStatus        int32             `gorm:"column:delivery_status;type:tinyint(1);comment:配送状态  1立即送出  0选择具体时间;not null;default:1;" json:"deliveryStatus"`      // 配送状态  1立即送出  0选择具体时间
	DeliveryTime          jsontime.JSONTime `gorm:"column:delivery_time;type:datetime;comment:送达时间;" json:"deliveryTime"`                                               // 送达时间
	PackAmount            int32             `gorm:"column:pack_amount;type:int;comment:打包费;" json:"packAmount"`                                                         // 打包费
	TablewareNumber       int32             `gorm:"column:tableware_number;type:int;comment:餐具数量;" json:"tablewareNumber"`                                              // 餐具数量
	TablewareStatus       int32             `gorm:"column:tableware_status;type:tinyint(1);comment:餐具数量状态  1按餐量提供  0选择具体数量;not null;default:1;" json:"tablewareStatus"` // 餐具数量状态  1按餐量提供  0选择具体数量
}

func (Orders) TableName() string {
	return "orders"
}

type OrderDetail struct {
	Id         int64   `gorm:"column:id;type:bigint;comment:主键;primaryKey;" json:"id"`               // 主键
	Name       string  `gorm:"column:name;type:varchar(32);comment:名字;" json:"name"`                 // 名字
	Pic        string  `gorm:"column:pic;type:longtext;comment:图片;" json:"pic"`                      // 图片
	OrderId    int64   `gorm:"column:order_id;type:bigint;comment:订单id;not null;" json:"orderId"`    // 订单id
	DishId     int64   `gorm:"column:dish_id;type:bigint;comment:菜品id;" json:"dishId"`               // 菜品id
	SetmealId  int64   `gorm:"column:setmeal_id;type:bigint;comment:套餐id;" json:"setmealId"`         // 套餐id
	DishFlavor string  `gorm:"column:dish_flavor;type:varchar(50);comment:口味;" json:"dishFlavor"`    // 口味
	Number     int32   `gorm:"column:number;type:int;comment:数量;not null;default:1;" json:"number"`  // 数量
	Amount     float64 `gorm:"column:amount;type:decimal(10, 2);comment:金额;not null;" json:"amount"` // 金额
}

func (OrderDetail) TableName() string {
	return "order_detail"
}
