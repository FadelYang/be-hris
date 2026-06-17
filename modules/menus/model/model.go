package model

import "github.com/google/uuid"

type Menu struct {
	ID           uuid.UUID  `json:"id"`
	Name         string     `json:"name"`
	Slug         string     `json:"slug"`
	ParentMenuID *uuid.UUID `json:"parent_menu_id,omitempty"`
}
