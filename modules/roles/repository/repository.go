package repository

import (
	"context"
	"errors"
	"net/http"
	"project-root/modules/roles/dto"
	"project-root/modules/roles/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RoleRepository interface {
	GetAll(ctx context.Context) (data []model.Role, httpCode int, err error)
	Create(ctx context.Context, form dto.CreateRole) (httpCode int, err error)
	GetByID(ctx context.Context, ID uuid.UUID) (data *model.Role, httpCode int, err error)
	UpdateByID(ctx context.Context, ID uuid.UUID, form dto.UpdateRole) (httpCode int, err error)
	DeleteByID(ctx context.Context, ID uuid.UUID) (httpCode int, err error)
}

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{
		db: db,
	}
}

func (r roleRepository) GetAll(ctx context.Context) (data []model.Role, httpCode int, err error) {
	if err := r.db.
		WithContext(ctx).
		Raw(qGet).
		Scan(&data).
		Error; err != nil {
		return nil, http.StatusBadRequest, err
	}

	return data, http.StatusOK, nil
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
	if err := r.db.
		WithContext(ctx).
		Raw(
			qGetByID,
			ID,
		).Scan(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, http.StatusNotFound, err
		}
		return nil, http.StatusBadRequest, err
	}

	return data, http.StatusOK, nil
}

func (r roleRepository) UpdateByID(ctx context.Context, ID uuid.UUID, form dto.UpdateRole) (httpCode int, err error) {
	if err := r.db.
		WithContext(ctx).
		Raw(
			qUpdateByID,
			form.Name,
			ID,
		).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return http.StatusNotFound, err
		}
		return http.StatusBadRequest, err
	}

	return http.StatusOK, nil
}

func (r roleRepository) DeleteByID(ctx context.Context, ID uuid.UUID) (httpCode int, err error) {
	if err := r.db.
		WithContext(ctx).
		Exec(
			qDeletebyID,
			ID,
		).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return http.StatusNotFound, err
		}
		return http.StatusBadRequest, err
	}

	return http.StatusOK, nil
}
