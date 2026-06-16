package common

import "project-root/tools"

type BaseResponse[T any] struct {
	Status     int               `json:"status"`
	Message    string            `json:"message"`
	Data       T                 `json:"data"`
	Pagination *tools.Pagination `json:"pagination,omitempty"`
}
