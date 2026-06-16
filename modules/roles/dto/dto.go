package dto

import "project-root/tools"

type CreateRole struct {
	Name string `json:"name"`
}

type UpdateRole struct {
	CreateRole
}

type Filter struct {
	Pagination tools.Pagination
}
