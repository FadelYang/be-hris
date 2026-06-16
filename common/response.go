package common

type Pagination struct {
	Page      int `json:"page"`
	Limit     int `json:"limit"`
	Offset    int `json:"offset"`
	TotalData int `json:"total_data"`
	TotalPage int `json:"total_page"`
}

type BaseResponse[T any] struct {
	Status     int         `json:"status"`
	Message    string      `json:"message"`
	Data       T           `json:"data"`
	Pagination *Pagination `json:"pagination,omitempty"`
}
