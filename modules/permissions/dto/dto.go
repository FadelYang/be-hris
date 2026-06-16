package dto

import "project-root/tools"

type Createpermission struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

type Updatepermission struct {
	Createpermission
}

type Filter struct {
	Pagination tools.Pagination
}
