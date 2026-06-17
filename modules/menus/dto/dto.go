package dto

import (
	"project-root/tools"

	"github.com/google/uuid"
)

type CreateMenu struct {
	Name         string     `json:"name"`
	Slug         string     `json:"-"`
	ParentMenuID *uuid.UUID `json:"parent_menu_id,omitempty"`
}

type UpdateMenu struct {
	CreateMenu
}

type Filter struct {
	Pagination tools.Pagination
}
