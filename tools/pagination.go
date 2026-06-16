package tools

import (
	"project-root/common"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Pagination struct {
	Page  int
	Limit int
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

func (p Pagination) Offset() int {
	return (p.Page - 1) * p.Limit
}

func GetPaginationQuery(c *gin.Context) common.Pagination {
	paginationQuery := common.Pagination{
		Page:  1,
		Limit: 10,
	}

	paginationQuery.Page, _ = strconv.Atoi(c.Query("page"))
	paginationQuery.Limit, _ = strconv.Atoi(c.Query("limit"))
	paginationQuery.Offset = (paginationQuery.Page - 1) * paginationQuery.Limit

	return paginationQuery
}
