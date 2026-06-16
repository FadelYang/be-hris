package tools

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Pagination struct {
	Page      int `json:"page"`
	Limit     int `json:"limit"`
	Offset    int `json:"offset"`
	TotalData int `json:"total_data"`
	TotalPage int `json:"total_page"`
}

func New(page, limit int) Pagination {
	if page <= 0 {
		page = 1
	}

	if limit <= 0 {
		limit = 10
	}

	if limit > 100 {
		limit = 100
	}

	return Pagination{
		Page:  page,
		Limit: limit,
	}
}

func (p Pagination) CalculateOffset() int {
	return (p.Page - 1) * p.Limit
}

func GetPaginationQuery(c *gin.Context) Pagination {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))

	p := New(page, limit)

	return Pagination{
		Page:   p.Page,
		Limit:  p.Limit,
		Offset: p.CalculateOffset(),
	}
}

func ConstructPaginationResponse(paginate Pagination) Pagination {
	pagination := Pagination{
		Page:      paginate.Page,
		Limit:     paginate.Limit,
		Offset:    paginate.Offset,
		TotalData: paginate.TotalData,
	}
	pagination.TotalPage = (pagination.TotalData + pagination.Limit - 1) / pagination.Limit

	return pagination
}
