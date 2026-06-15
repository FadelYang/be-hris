package providers

import (
	"project-root/modules/roles/handler"
	"project-root/modules/roles/repository"
	"project-root/modules/roles/usecase"

	"gorm.io/gorm"
)

type Provider struct {
	RoleHandler *handler.RoleHandler
}

func NewProvider(db *gorm.DB) *Provider {
	repo := repository.NewRoleRepository(db)
	usecase := usecase.NewRoleUsecase(repo)
	handler := handler.NewRoleHandler(usecase)

	return &Provider{
		RoleHandler: handler,
	}
}
