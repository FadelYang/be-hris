package repository

import (
	"context"
	"net/http"
	"project-root/modules/permissions/dto"
	"project-root/modules/permissions/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PermissionRepository interface {
	GetAll(ctx context.Context, filter dto.Filter) (data []model.Permission, totalData int, httpCode int, err error)
	Create(ctx context.Context, form dto.Createpermission) (httpCode int, err error)
	GetByID(ctx context.Context, ID uuid.UUID) (data *model.Permission, httpCode int, err error)
	UpdateByID(ctx context.Context, ID uuid.UUID, form dto.Updatepermission) (httpCode int, err error)
	DeleteByID(ctx context.Context, ID uuid.UUID) (httpCode int, err error)
}

type permissionRepository struct {
	db *gorm.DB
}

func NewPermissionRepository(db *gorm.DB) PermissionRepository {
	return &permissionRepository{
		db: db,
	}
}

func (r permissionRepository) GetAll(ctx context.Context, filter dto.Filter) (data []model.Permission, totalData int, httpCode int, err error) {
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

func (r permissionRepository) Create(ctx context.Context, form dto.Createpermission) (httpCode int, err error) {
	if err := r.db.
		WithContext(ctx).
		Exec(
			qCreate,
			form.Name,
			form.Description,
		).Error; err != nil {
		return http.StatusBadRequest, err
	}

	return http.StatusOK, nil
}

func (r permissionRepository) GetByID(ctx context.Context, ID uuid.UUID) (data *model.Permission, httpCode int, err error) {
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

func (r permissionRepository) UpdateByID(ctx context.Context, ID uuid.UUID, form dto.Updatepermission) (httpCode int, err error) {
	tx := r.db.
		WithContext(ctx).
		Exec(
			qUpdateByID,
			form.Name,
			form.Description,
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

func (r permissionRepository) DeleteByID(ctx context.Context, ID uuid.UUID) (httpCode int, err error) {
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
