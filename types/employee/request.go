package employee

type RegisterRequest struct {
	Password string `json:"password"`
	Account  string `json:"account"`
}

type LoginRequest struct {
	Password string `json:"password"`
	Account  string `json:"account"`
}

type EditPasswordRequest struct {
	OldPwd string `json:"oldPwd"`
	NewPwd string `json:"newPwd"`
}

type AddEditRequest struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Account  string `json:"account"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Age      int    `json:"age"`
	Gender   int    `json:"gender"`
	Pic      string `json:"pic"`
}

type PageRequest struct {
	Name     string `json:"name,omitempty"`
	Page     int    `json:"page,omitempty"`
	PageSize int    `json:"pageSize,omitempty"`
}
