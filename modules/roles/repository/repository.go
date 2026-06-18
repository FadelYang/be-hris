package repository

import (
	"context"
	"fmt"
	"net/http"
	"project-root/modules/roles/dto"
	"project-root/modules/roles/model"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RoleRepository interface {
	GetAll(ctx context.Context, filter dto.Filter) (data []model.Role, totalData int, httpCode int, err error)
	Create(ctx context.Context, form dto.CreateRole) (httpCode int, err error)
	GetByID(ctx context.Context, ID uuid.UUID) (data *model.Role, httpCode int, err error)
	UpdateByID(ctx context.Context, ID uuid.UUID, form dto.UpdateRole) (httpCode int, err error)
	DeleteByID(ctx context.Context, ID uuid.UUID) (httpCode int, err error)
	AssignMenusPermissions(ctx context.Context, roleID uuid.UUID, form dto.AssignMenusPermissions) (httpCode int, err error)
}

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{
		db: db,
	}
}

func (r roleRepository) GetAll(ctx context.Context, filter dto.Filter) (data []model.Role, totalData int, httpCode int, err error) {
	if err := r.db.
		WithContext(ctx).
		Raw(
			qGet,
			filter.Pagination.Limit,
			filter.Pagination.Offset,
		).
		Scan(&data).
		Error; err != nil {
		return nil, totalData, http.StatusBadRequest, err
	}

	if err := r.db.
		WithContext(ctx).
		Raw(qCount).
		Scan(&totalData).
		Error; err != nil {
		return nil, totalData, http.StatusBadRequest, err
	}

	return data, totalData, http.StatusOK, nil
}

func (r roleRepository) Create(ctx context.Context, form dto.CreateRole) (httpCode int, err error) {
	if err := r.db.
		WithContext(ctx).
		Exec(
			qCreate,
			form.Name,
		).Error; err != nil {
		return http.StatusBadRequest, err
	}

	return http.StatusOK, nil
}

func (r roleRepository) GetByID(ctx context.Context, ID uuid.UUID) (data *model.Role, httpCode int, err error) {
	tx := r.db.
		WithContext(ctx).
		Raw(
			qGetByID,
			ID,
		).Scan(&data)

	if tx.Error != nil {
		return nil, http.StatusBadRequest, tx.Error
	}

	if tx.RowsAffected == 0 {
		return nil, http.StatusNotFound, gorm.ErrRecordNotFound
	}

	return data, http.StatusOK, nil
}

func (r roleRepository) UpdateByID(ctx context.Context, ID uuid.UUID, form dto.UpdateRole) (httpCode int, err error) {
	tx := r.db.
		WithContext(ctx).
		Exec(
			qUpdateByID,
			form.Name,
			ID,
		)

	if tx.Error != nil {
		return http.StatusBadRequest, tx.Error
	}

	if tx.RowsAffected == 0 {
		return http.StatusNotFound, gorm.ErrRecordNotFound
	}

	return http.StatusOK, nil
}

func (r roleRepository) DeleteByID(ctx context.Context, ID uuid.UUID) (httpCode int, err error) {
	tx := r.db.
		WithContext(ctx).
		Exec(
			qDeletebyID,
			ID,
		)

	if tx.Error != nil {
		return http.StatusBadRequest, tx.Error
	}

	if tx.RowsAffected == 0 {
		return http.StatusNotFound, gorm.ErrRecordNotFound
	}

	return http.StatusOK, nil
}

func (r roleRepository) constructAssignMenusPermissionsValues(
	roleID uuid.UUID,
	data dto.AssignMenusPermissions,
) []string {

	values := []string{}

	for _, item := range data.Items {
		values = append(
			values,
			fmt.Sprintf(
				"('%s','%s','%s')",
				roleID,
				item.MenuID,
				item.PermissionID,
			),
		)
	}

	return values
}

func (r roleRepository) AssignMenusPermissions(
	ctx context.Context,
	roleID uuid.UUID,
	form dto.AssignMenusPermissions,
) (httpCode int, err error) {

	values := r.constructAssignMenusPermissionsValues(
		roleID,
		form,
	)

	tx := r.db.
		WithContext(ctx).
		Begin()

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		}
	}()

	if tx.Error != nil {
		return http.StatusInternalServerError, tx.Error
	}

	if err := tx.Exec(
		qDeleteAssignedMenusPermissions,
		roleID,
	).Error; err != nil {
		tx.Rollback()
		return http.StatusBadRequest, err
	}

	if len(values) > 0 {
		if err := tx.Exec(
			qAssignMenusPermissions +
				strings.Join(values, ","),
		).Error; err != nil {
			tx.Rollback()
			return http.StatusBadRequest, err
		}
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
