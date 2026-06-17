package handler

import (
	"context"
	"net/http"
	"project-root/common"
	"project-root/modules/menus/dto"
	"project-root/modules/menus/model"
	"project-root/modules/menus/usecase"
	"project-root/tools"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type MenuHandler struct {
	menuUsecase usecase.MenuUsecase
}

func NewmenuHandler(menuUsecase usecase.MenuUsecase) *MenuHandler {
	return &MenuHandler{
		menuUsecase: menuUsecase,
	}
}

// @Tags 					Menus
// @Summary				Get menus
// @Description 	Get all menus
// @Accept 				json
// @Produce 			json
// @Success				200 {object} common.BaseResponse[[]model.Menu]
// @Param         page query int false "Page number" default(1)
// @Param         limit query int false "Items per page" default(10)
// @Router				/menus [get]
func (h *MenuHandler) GetAll(c *gin.Context) {
	paginate := tools.GetPaginationQuery(c)

	filter := dto.Filter{
		Pagination: paginate,
	}

	data, pagination, httpCode, err := h.menuUsecase.GetAll(c.Request.Context(), filter)
	if err != nil {
		tools.HandleLogError(err, "")
		tools.HandlerSimpleError(c, httpCode, "", err)
		return
	}

	c.JSON(
		httpCode,
		common.BaseResponse[[]model.Menu]{
			Status:     httpCode,
			Message:    "succes get menus data",
			Data:       data,
			Pagination: &pagination,
		},
	)
}

// @Tags 					Menus
// @Summary				Get Menu Detail By ID
// @Description 	Get menu detail by id
// @Accept 				json
// @Produce 			json
// @Success				200 {object} common.BaseResponse[model.Menu]
// @Param         id path string true "menu ID"
// @Router				/menus/{id} [get]
func (h *MenuHandler) GetByID(c *gin.Context) {
	MenuID := c.Param("id")
	parsedMenuID, err := uuid.Parse(MenuID)
	if err != nil {
		tools.HandleLogError(err, "")
		tools.HandlerSimpleError(c, http.StatusBadRequest, "", err)
		return
	}

	data, httpCode, err := h.menuUsecase.GetByID(context.Background(), parsedMenuID)
	if err != nil {
		tools.HandleLogError(err, "")
		tools.HandlerSimpleError(c, httpCode, "", err)
		return
	}

	c.JSON(
		httpCode,
		common.BaseResponse[model.Menu]{
			Status:  httpCode,
			Message: "successfully get menu detail",
			Data:    *data,
		},
	)
}

// @Tags 					Menus
// @Summary				Create Menu
// @Description 	Create a menu
// @Accept 				json
// @Produce 			json
// @Success				201 {object} common.BaseResponse[any]
// @Router				/menus [post]
// @Param					request body dto.CreateMenu true "request body for create a menu [RAW]"
func (h *MenuHandler) Create(c *gin.Context) {
	var form dto.CreateMenu
	if err := c.ShouldBindBodyWithJSON(&form); err != nil {
		tools.HandleLogError(err, "")
		tools.HandlerSimpleError(c, http.StatusBadRequest, "", err)
	}

	form.Slug = tools.GenerateSlug(form.Name)

	httpCode, err := h.menuUsecase.Create(c.Request.Context(), form)
	if err != nil {
		tools.HandleLogError(err, "")
		tools.HandlerSimpleError(c, httpCode, "", err)
		return
	}

	c.JSON(
		httpCode,
		common.BaseResponse[any]{
			Status:  httpCode,
			Message: "successfully create a menu",
			Data:    nil,
		},
	)
}

// @Tags 					Menus
// @Summary				Update Menu
// @Description 	Update a menu
// @Accept 				json
// @Produce 			json
// @Success				201 {object} common.BaseResponse[any]
// @Router				/menus/{id} [put]
// @Param         id path string true "Menu ID"
// @Param					request body dto.UpdateMenu true "request body for update a menu [RAW]"
func (h *MenuHandler) UpdateByID(c *gin.Context) {
	var form dto.UpdateMenu
	if err := c.ShouldBindBodyWithJSON(&form); err != nil {
		tools.HandleLogError(err, "")
		tools.HandlerSimpleError(c, http.StatusBadRequest, "", err)
	}

	menuID := c.Param("id")
	parsedMenuID, err := uuid.Parse(menuID)
	if err != nil {
		tools.HandleLogError(err, "")
		tools.HandlerSimpleError(c, http.StatusBadRequest, "", err)
		return
	}

	httpCode, err := h.menuUsecase.UpdateByID(c.Request.Context(), parsedMenuID, form)
	if err != nil {
		tools.HandleLogError(err, "")
		tools.HandlerSimpleError(c, httpCode, "", err)
		return
	}

	c.JSON(
		httpCode,
		common.BaseResponse[any]{
			Status:  httpCode,
			Message: "successfully update a menu",
			Data:    nil,
		},
	)
}

// @Tags 					Menus
// @Summary				Delete Menu By ID
// @Description 	Delete menu by its id
// @Accept 				json
// @Produce 			json
// @Success				200 {object} common.BaseResponse[any]
// @Param         id path string true "menu ID"
// @Router				/menus/{id} [delete]
func (h *MenuHandler) DeleteByID(c *gin.Context) {
	menuID := c.Param("id")
	parsedMenuID, err := uuid.Parse(menuID)
	if err != nil {
		tools.HandleLogError(err, "")
		tools.HandlerSimpleError(c, http.StatusBadRequest, "", err)
		return
	}

	httpCode, err := h.menuUsecase.DeleteByID(context.Background(), parsedMenuID)
	if err != nil {
		tools.HandleLogError(err, "")
		tools.HandlerSimpleError(c, httpCode, "", err)
		return
	}

	c.JSON(
		httpCode,
		common.BaseResponse[any]{
			Status:  httpCode,
			Message: "successfully delete menu",
			Data:    nil,
		},
	)
}
