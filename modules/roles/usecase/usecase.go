package usecase

import (
	"context"
	"project-root/modules/roles/dto"
	"project-root/modules/roles/model"
	"project-root/modules/roles/repository"
	"project-root/tools"

	"github.com/google/uuid"
)

type RoleUsecase interface {
	GetAll(ctx context.Context, filter dto.Filter) (data []model.Role, pagination tools.Pagination, httpCode int, err error)
	Create(ctx context.Context, form dto.CreateRole) (httpCode int, err error)
	GetByID(ctx context.Context, ID uuid.UUID) (data *model.Role, httpCode int, err error)
	UpdateByID(ctx context.Context, ID uuid.UUID, form dto.UpdateRole) (httpCode int, err error)
	DeleteByID(ctx context.Context, ID uuid.UUID) (httpCode int, err error)
	AssignMenusPermissions(ctx context.Context, roleID uuid.UUID, form dto.AssignMenusPermissions) (httpCode int, err error)
}

type roleUsecase struct {
	roleRepository repository.RoleRepository
}

func NewRoleUsecase(roleRepository repository.RoleRepository) RoleUsecase {
	return &roleUsecase{
		roleRepository: roleRepository,
	}
}

func (u roleUsecase) GetAll(ctx context.Context, filter dto.Filter) (data []model.Role, pagination tools.Pagination, httpCode int, err error) {
	data, total, httpCode, err := u.roleRepository.GetAll(ctx, filter)
	filter.Pagination.TotalData = total

	pagination = tools.ConstructPaginationResponse(filter.Pagination)

	return data, pagination, httpCode, err
}

func (u roleUsecase) Create(ctx context.Context, form dto.CreateRole) (httpCode int, err error) {
	httpCode, err = u.roleRepository.Create(ctx, form)

	return httpCode, err
}

func (u roleUsecase) GetByID(ctx context.Context, ID uuid.UUID) (data *model.Role, httpCode int, err error) {
	data, httpCode, err = u.roleRepository.GetByID(ctx, ID)

	return data, httpCode, err
}

func (u roleUsecase) UpdateByID(ctx context.Context, ID uuid.UUID, form dto.UpdateRole) (httpCode int, err error) {
	httpCode, err = u.roleRepository.UpdateByID(ctx, ID, form)

	return httpCode, err
}

func (u roleUsecase) DeleteByID(ctx context.Context, ID uuid.UUID) (httpCode int, err error) {
	httpCode, err = u.roleRepository.DeleteByID(ctx, ID)

	return httpCode, err
}

func (u roleUsecase) AssignMenusPermissions(ctx context.Context, roleID uuid.UUID, form dto.AssignMenusPermissions) (httpCode int, err error) {
	httpCode, err = u.roleRepository.AssignMenusPermissions(ctx, roleID, form)

	return httpCode, err
}
