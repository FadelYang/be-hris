package dto

import (
	"project-root/tools"

	"github.com/google/uuid"
)

type CreateRole struct {
	Name string `json:"name"`
}

type UpdateRole struct {
	CreateRole
}

type AssignMenusPermissions struct {
	MenuIDs       []uuid.UUID `json:"menu_ids"`
	PermissionIDs []uuid.UUID `json:"permission_ids"`
}

type Filter struct {
	Pagination tools.Pagination
}
