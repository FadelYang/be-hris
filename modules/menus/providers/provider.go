package providers

import (
	"project-root/modules/menus/handler"
	"project-root/modules/menus/repository"
	"project-root/modules/menus/usecase"

	"gorm.io/gorm"
)

type Provider struct {
	MenuHandler *handler.MenuHandler
}

func NewProvider(db *gorm.DB) *Provider {
	repo := repository.NewMenuRepository(db)
	usecase := usecase.NewMenuUsecase(repo)
	handler := handler.NewmenuHandler(usecase)

	return &Provider{
		MenuHandler: handler,
	}
}
