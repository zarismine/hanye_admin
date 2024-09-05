package order

type PageReq struct {
	BeginTime string `json:"beginTime,omitempty"`
	EndTime   string `json:"endTime,omitempty"`
	Number    string `json:"number,omitempty"`
	Page      int    `json:"page,omitempty"`
	PageSize  int    `json:"pageSize,omitempty"`
	Phone     string `json:"phone,omitempty"`
	Status    int    `json:"status,omitempty"`
	UserID    int    `json:"userId,omitempty"`
}

type ConfirmReq struct {
	Id     int `json:"id"`
	Status int `json:"status"`
}

type RejectReq struct {
	Id              int    `json:"id"`
	RejectionReason string `json:"rejectionReason"`
}
