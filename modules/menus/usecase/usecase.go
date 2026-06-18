package usecase

import (
	"context"
	"project-root/modules/menus/dto"
	"project-root/modules/menus/model"
	"project-root/modules/menus/repository"
	"project-root/tools"

	"github.com/google/uuid"
)

type MenuUsecase interface {
	GetAll(ctx context.Context, filter dto.Filter) (data []model.Menu, pagination tools.Pagination, httpCode int, err error)
	Create(ctx context.Context, form dto.CreateMenu) (httpCode int, err error)
	GetByID(ctx context.Context, ID uuid.UUID) (data *model.Menu, httpCode int, err error)
	UpdateByID(ctx context.Context, ID uuid.UUID, form dto.UpdateMenu) (httpCode int, err error)
	DeleteByID(ctx context.Context, ID uuid.UUID) (httpCode int, err error)
}

type menuUsecase struct {
	menuRepository repository.MenuRepository
}

func NewMenuUsecase(menuRepository repository.MenuRepository) MenuUsecase {
	return &menuUsecase{
		menuRepository: menuRepository,
	}
}

func (u *menuUsecase) GetAll(ctx context.Context, filter dto.Filter) (data []model.Menu, pagination tools.Pagination, httpCode int, err error) {
	data, total, httpCode, err := u.menuRepository.GetAll(ctx, filter)
	filter.Pagination.TotalData = total

	pagination = tools.ConstructPaginationResponse(filter.Pagination)

	return data, pagination, httpCode, err
}

func (u *menuUsecase) Create(ctx context.Context, form dto.CreateMenu) (httpCode int, err error) {
	httpCode, err = u.menuRepository.Create(ctx, form)

	return httpCode, err
}

func (u *menuUsecase) GetByID(ctx context.Context, ID uuid.UUID) (data *model.Menu, httpCode int, err error) {
	data, httpCode, err = u.menuRepository.GetByID(ctx, ID)

	return data, httpCode, err
}

func (u *menuUsecase) UpdateByID(ctx context.Context, ID uuid.UUID, form dto.UpdateMenu) (httpCode int, err error) {
	httpCode, err = u.menuRepository.UpdateByID(ctx, ID, form)

	return httpCode, err
}

func (u *menuUsecase) DeleteByID(ctx context.Context, ID uuid.UUID) (httpCode int, err error) {
	httpCode, err = u.menuRepository.DeleteByID(ctx, ID)

	return httpCode, err
}
