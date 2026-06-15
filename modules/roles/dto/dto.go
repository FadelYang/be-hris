package dto

type CreateRole struct {
	Name string `json:"name"`
}

type UpdateRole struct {
	CreateRole
}
