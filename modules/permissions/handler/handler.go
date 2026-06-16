package handler

import (
	"context"
	"fmt"
	"net/http"
	"project-root/common"
	"project-root/modules/permissions/dto"
	"project-root/modules/permissions/model"
	"project-root/modules/permissions/usecase"
	"project-root/tools"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PermissionHandler struct {
	permissionUsecase usecase.PermissionUsecase
}

func NewPermissionHandler(permissionUsecase usecase.PermissionUsecase) *PermissionHandler {
	return &PermissionHandler{
		permissionUsecase: permissionUsecase,
	}
}

// @Tags 					permissions
// @Summary				Get permissions
// @Description 	Get all permissions
// @Accept 				json
// @Produce 			json
// @Success				200 {object} common.BaseResponse[[]model.Permission]
// @Param         page query int false "Page number" default(1)
// @Param         limit query int false "Items per page" default(10)
// @Router				/permissions [get]
func (h *PermissionHandler) GetAll(c *gin.Context) {
	paginate := tools.GetPaginationQuery(c)

	filter := dto.Filter{
		Pagination: paginate,
	}

	data, pagination, httpCode, err := h.permissionUsecase.GetAll(c.Request.Context(), filter)
	if err != nil {
		errMsg := "failed to get permissions data"
		tools.HandleLogError(err, errMsg)
		tools.HandlerSimpleError(c, httpCode, errMsg, err)
		return
	}

	c.JSON(
		httpCode,
		common.BaseResponse[[]model.Permission]{
			Status:     httpCode,
			Message:    "succes get permissions data",
			Data:       data,
			Pagination: &pagination,
		},
	)
}

// @Tags 					permissions
// @Summary				Get Permission Detail By ID
// @Description 	Get permission detail by id
// @Accept 				json
// @Produce 			json
// @Success				200 {object} common.BaseResponse[model.Permission]
// @Param         id path string true "permission ID"
// @Router				/permissions/{id} [get]
func (h *PermissionHandler) GetByID(c *gin.Context) {
	permissionID := c.Param("id")
	parsedpermissionID, err := uuid.Parse(permissionID)
	if err != nil {
		errMsg := fmt.Sprintf("failed to parsed %s as permissionID", permissionID)
		tools.HandleLogError(err, errMsg)
		tools.HandlerSimpleError(c, http.StatusBadRequest, errMsg, err)
		return
	}

	data, httpCode, err := h.permissionUsecase.GetByID(context.Background(), parsedpermissionID)
	if err != nil {
		tools.HandleLogError(err, "")
		tools.HandlerSimpleError(c, httpCode, "", err)
		return
	}

	c.JSON(
		httpCode,
		common.BaseResponse[model.Permission]{
			Status:  httpCode,
			Message: "successfully get permission detail",
			Data:    *data,
		},
	)
}

// @Tags 					permissions
// @Summary				Create permission
// @Description 	Create a permission
// @Accept 				json
// @Produce 			json
// @Success				201 {object} common.BaseResponse[any]
// @Router				/permissions [post]
// @Param					request body dto.Createpermission true "request body for create a permission [RAW]"
func (h *PermissionHandler) Create(c *gin.Context) {
	var form dto.Createpermission
	if err := c.ShouldBindBodyWithJSON(&form); err != nil {
		errMsg := "failed to create permission"
		tools.HandleLogError(err, errMsg)
		tools.HandlerSimpleError(c, http.StatusBadRequest, errMsg, err)
	}

	httpCode, err := h.permissionUsecase.Create(c.Request.Context(), form)
	if err != nil {
		errMsg := "failed to create a permission"
		tools.HandleLogError(err, errMsg)
		tools.HandlerSimpleError(c, httpCode, errMsg, err)
		return
	}

	c.JSON(
		httpCode,
		common.BaseResponse[any]{
			Status:  httpCode,
			Message: "successfully create a permission",
			Data:    nil,
		},
	)
}

// @Tags 					permissions
// @Summary				Update permission
// @Description 	Update a permission
// @Accept 				json
// @Produce 			json
// @Success				201 {object} common.BaseResponse[any]
// @Router				/permissions/{id} [put]
// @Param         id path string true "permission ID"
// @Param					request body dto.Updatepermission true "request body for update a permission [RAW]"
func (h *PermissionHandler) UpdateByID(c *gin.Context) {
	var form dto.Updatepermission
	if err := c.ShouldBindBodyWithJSON(&form); err != nil {
		errMsg := "failed to update permission"
		tools.HandleLogError(err, errMsg)
		tools.HandlerSimpleError(c, http.StatusBadRequest, errMsg, err)
	}

	permissionID := c.Param("id")
	parsedpermissionID, err := uuid.Parse(permissionID)
	if err != nil {
		errMsg := fmt.Sprintf("failed to parsed %s as permissionID", permissionID)
		tools.HandleLogError(err, errMsg)
		tools.HandlerSimpleError(c, http.StatusBadRequest, errMsg, err)
		return
	}

	httpCode, err := h.permissionUsecase.UpdateByID(c.Request.Context(), parsedpermissionID, form)
	if err != nil {
		errMsg := "failed to update a permission"
		tools.HandleLogError(err, errMsg)
		tools.HandlerSimpleError(c, httpCode, errMsg, err)
		return
	}

	c.JSON(
		httpCode,
		common.BaseResponse[any]{
			Status:  httpCode,
			Message: "successfully update a permission",
			Data:    nil,
		},
	)
}

// @Tags 					permissions
// @Summary				Delete permission By ID
// @Description 	Delete permission by its id
// @Accept 				json
// @Produce 			json
// @Success				200 {object} common.BaseResponse[any]
// @Param         id path string true "permission ID"
// @Router				/permissions/{id} [delete]
func (h *PermissionHandler) DeleteByID(c *gin.Context) {
	permissionID := c.Param("id")
	parsedpermissionID, err := uuid.Parse(permissionID)
	if err != nil {
		errMsg := fmt.Sprintf("failed to parsed %s as permissionID", permissionID)
		tools.HandleLogError(err, errMsg)
		tools.HandlerSimpleError(c, http.StatusBadRequest, errMsg, err)
		return
	}

	httpCode, err := h.permissionUsecase.DeleteByID(context.Background(), parsedpermissionID)
	if err != nil {
		errMsg := "failed to delete permission"
		tools.HandleLogError(err, errMsg)
		tools.HandlerSimpleError(c, httpCode, errMsg, err)
		return
	}

	c.JSON(
		httpCode,
		common.BaseResponse[any]{
			Status:  httpCode,
			Message: "successfully delete permission",
			Data:    nil,
		},
	)
}
