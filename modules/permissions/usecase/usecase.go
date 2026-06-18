package usecase

import (
	"context"
	"project-root/modules/permissions/dto"
	"project-root/modules/permissions/model"
	"project-root/modules/permissions/repository"
	"project-root/tools"

	"github.com/google/uuid"
)

type PermissionUsecase interface {
	GetAll(ctx context.Context, filter dto.Filter) (data []model.Permission, pagination tools.Pagination, httpCode int, err error)
	Create(ctx context.Context, form dto.Createpermission) (httpCode int, err error)
	GetByID(ctx context.Context, ID uuid.UUID) (data *model.Permission, httpCode int, err error)
	UpdateByID(ctx context.Context, ID uuid.UUID, form dto.Updatepermission) (httpCode int, err error)
	DeleteByID(ctx context.Context, ID uuid.UUID) (httpCode int, err error)
}

type permissionUsecase struct {
	permissionRepository repository.PermissionRepository
}

func NewPermissionUsecase(permissionRepository repository.PermissionRepository) PermissionUsecase {
	return &permissionUsecase{
		permissionRepository: permissionRepository,
	}
}

func (u *permissionUsecase) GetAll(ctx context.Context, filter dto.Filter) (data []model.Permission, pagination tools.Pagination, httpCode int, err error) {
	data, total, httpCode, err := u.permissionRepository.GetAll(ctx, filter)
	filter.Pagination.TotalData = total

	pagination = tools.ConstructPaginationResponse(filter.Pagination)

	return data, pagination, httpCode, err
}

func (u *permissionUsecase) Create(ctx context.Context, form dto.Createpermission) (httpCode int, err error) {
	httpCode, err = u.permissionRepository.Create(ctx, form)

	return httpCode, err
}

func (u *permissionUsecase) GetByID(ctx context.Context, ID uuid.UUID) (data *model.Permission, httpCode int, err error) {
	data, httpCode, err = u.permissionRepository.GetByID(ctx, ID)

	return data, httpCode, err
}

func (u *permissionUsecase) UpdateByID(ctx context.Context, ID uuid.UUID, form dto.Updatepermission) (httpCode int, err error) {
	httpCode, err = u.permissionRepository.UpdateByID(ctx, ID, form)

	return httpCode, err
}

func (u *permissionUsecase) DeleteByID(ctx context.Context, ID uuid.UUID) (httpCode int, err error) {
	httpCode, err = u.permissionRepository.DeleteByID(ctx, ID)

	return httpCode, err
}
