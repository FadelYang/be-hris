package providers

import (
	"project-root/modules/permissions/handler"
	"project-root/modules/permissions/repository"
	"project-root/modules/permissions/usecase"

	"gorm.io/gorm"
)

type Provider struct {
	PermissionHandler *handler.PermissionHandler
}

func NewProvider(db *gorm.DB) *Provider {
	repo := repository.NewPermissionRepository(db)
	usecase := usecase.NewPermissionUsecase(repo)
	handler := handler.NewPermissionHandler(usecase)

	return &Provider{
		PermissionHandler: handler,
	}
}
