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

type MenuPermission struct {
	MenuID       uuid.UUID `json:"menu_id"`
	PermissionID uuid.UUID `json:"permission_id"`
}

type AssignMenusPermissions struct {
	Items []MenuPermission `json:"menu_permission"`
}

type Filter struct {
	Pagination tools.Pagination
}
