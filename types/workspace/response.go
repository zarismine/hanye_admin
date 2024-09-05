package workspace

type Response struct {
	CountUp   int64 `json:"sold"`
	CountDown int64 `json:"discontinued"`
}

type OrderResponse struct {
	WaitingOrders   int64 `json:"waitingOrders"`
	DeliveredOrders int64 `json:"deliveredOrders"`
	CompletedOrders int64 `json:"completedOrders"`
	CancelledOrders int64 `json:"cancelledOrders"`
	AllOrders       int64 `json:"allOrders"`
}

type BusinessResponse struct {
	Turnover            float64 `json:"turnover"`
	ValidOrderCount     int64   `json:"validOrderCount"`
	OrderCompletionRate float64 `json:"orderCompletionRate"`
	UnitPrice           float64 `json:"unitPrice"`
	NewUsers            int64   `json:"newUsers"`
}
