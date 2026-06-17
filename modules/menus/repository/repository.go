package repository

import (
	"context"
	"net/http"
	"project-root/modules/menus/dto"
	"project-root/modules/menus/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MenuRepository interface {
	GetAll(ctx context.Context, filter dto.Filter) (data []model.Menu, totalData int, httpCode int, err error)
	Create(ctx context.Context, form dto.CreateMenu) (httpCode int, err error)
	GetByID(ctx context.Context, ID uuid.UUID) (data *model.Menu, httpCode int, err error)
	UpdateByID(ctx context.Context, ID uuid.UUID, form dto.UpdateMenu) (httpCode int, err error)
	DeleteByID(ctx context.Context, ID uuid.UUID) (httpCode int, err error)
}

type menuRepository struct {
	db *gorm.DB
}

func NewMenuRepository(db *gorm.DB) MenuRepository {
	return &menuRepository{
		db: db,
	}
}

func (r menuRepository) GetAll(ctx context.Context, filter dto.Filter) (data []model.Menu, totalData int, httpCode int, err error) {
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

func (r menuRepository) Create(ctx context.Context, form dto.CreateMenu) (httpCode int, err error) {
	if err := r.db.
		WithContext(ctx).
		Exec(
			qCreate,
			form.Name,
			form.Slug,
			form.ParentMenuID,
		).Error; err != nil {
		return http.StatusBadRequest, err
	}

	return http.StatusOK, nil
}

func (r menuRepository) GetByID(ctx context.Context, ID uuid.UUID) (data *model.Menu, httpCode int, err error) {
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

func (r menuRepository) UpdateByID(ctx context.Context, ID uuid.UUID, form dto.UpdateMenu) (httpCode int, err error) {
	tx := r.db.
		WithContext(ctx).
		Exec(
			qUpdateByID,
			form.Name,
			form.Slug,
			form.ParentMenuID,
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

func (r menuRepository) DeleteByID(ctx context.Context, ID uuid.UUID) (httpCode int, err error) {
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
