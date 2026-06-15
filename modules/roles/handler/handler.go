package handler

import (
	"context"
	"fmt"
	"net/http"
	"project-root/common"
	"project-root/modules/roles/dto"
	"project-root/modules/roles/model"
	"project-root/modules/roles/usecase"
	"project-root/tools"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type roleHandler struct {
	roleUsecase usecase.RoleUsecase
}

func NewRoleHandler(role usecase.RoleUsecase) *roleHandler {
	return &roleHandler{
		roleUsecase: role,
	}
}

// @Tags 					roles
// @Summary				Get Roles
// @Description 	Get all roles
// @Accept 				json
// @Produce 			json
// @Success				200 {object} common.BaseResponse[[]model.Role]
// @Router				/roles [get]
func (h *roleHandler) GetAll(c *gin.Context) {
	data, httpCode, err := h.roleUsecase.GetAll(c.Request.Context())
	if err != nil {
		errMsg := "failed to get roles data"
		tools.HandleLogError(err, errMsg)
		tools.HandlerSimpleError(c, httpCode, errMsg, err)
		return
	}

	c.JSON(
		httpCode,
		common.BaseResponse[[]model.Role]{
			Status:  httpCode,
			Message: "succes get roles data",
			Data:    data,
		},
	)
}

// @Tags 					roles
// @Summary				Get Role Detail By ID
// @Description 	Get role detail by id
// @Accept 				json
// @Produce 			json
// @Success				200 {object} common.BaseResponse[model.Role]
// @Param         id path string true "Role ID"
// @Router				/roles/{id} [get]
func (h *roleHandler) GetByID(c *gin.Context) {
	roleID := c.Param("id")
	parsedRoleID, err := uuid.Parse(roleID)
	if err != nil {
		errMsg := fmt.Sprintf("failed to parsed %s as roleID", roleID)
		tools.HandleLogError(err, errMsg)
		tools.HandlerSimpleError(c, http.StatusBadRequest, errMsg, err)
		return
	}

	data, httpCode, err := h.roleUsecase.GetByID(context.Background(), parsedRoleID)
	if err != nil {
		errMsg := "failed to get role detail"
		tools.HandleLogError(err, errMsg)
		tools.HandlerSimpleError(c, httpCode, errMsg, err)
		return
	}

	c.JSON(
		httpCode,
		common.BaseResponse[model.Role]{
			Status:  httpCode,
			Message: "successfully get role detail",
			Data:    *data,
		},
	)
}

// @Tags 					roles
// @Summary				Create Role
// @Description 	Create a role
// @Accept 				json
// @Produce 			json
// @Success				201 {object} common.BaseResponse[any]
// @Router				/roles [post]
// @Param					request body dto.CreateRole true "request body for create a role [RAW]"
func (h *roleHandler) Create(c *gin.Context) {
	var form dto.CreateRole
	if err := c.ShouldBindBodyWithJSON(&form); err != nil {
		errMsg := "failed to create role"
		tools.HandleLogError(err, errMsg)
		tools.HandlerSimpleError(c, http.StatusBadRequest, errMsg, err)
	}

	httpCode, err := h.roleUsecase.Create(c.Request.Context(), form)
	if err != nil {
		errMsg := "failed to create a role"
		tools.HandleLogError(err, errMsg)
		tools.HandlerSimpleError(c, httpCode, errMsg, err)
		return
	}

	c.JSON(
		httpCode,
		common.BaseResponse[any]{
			Status:  httpCode,
			Message: "successfully create a role",
			Data:    nil,
		},
	)
}

// @Tags 					roles
// @Summary				Update Role
// @Description 	Update a role
// @Accept 				json
// @Produce 			json
// @Success				201 {object} common.BaseResponse[any]
// @Router				/roles/{id} [put]
// @Param         id path string true "Role ID"
// @Param					request body dto.UpdateRole true "request body for update a role [RAW]"
func (h *roleHandler) UpdateByID(c *gin.Context) {
	var form dto.UpdateRole
	if err := c.ShouldBindBodyWithJSON(&form); err != nil {
		errMsg := "failed to update role"
		tools.HandleLogError(err, errMsg)
		tools.HandlerSimpleError(c, http.StatusBadRequest, errMsg, err)
	}

	roleID := c.Param("id")
	parsedRoleID, err := uuid.Parse(roleID)
	if err != nil {
		errMsg := fmt.Sprintf("failed to parsed %s as roleID", roleID)
		tools.HandleLogError(err, errMsg)
		tools.HandlerSimpleError(c, http.StatusBadRequest, errMsg, err)
		return
	}

	httpCode, err := h.roleUsecase.UpdateByID(c.Request.Context(), parsedRoleID, form)
	if err != nil {
		errMsg := "failed to update a role"
		tools.HandleLogError(err, errMsg)
		tools.HandlerSimpleError(c, httpCode, errMsg, err)
		return
	}

	c.JSON(
		httpCode,
		common.BaseResponse[any]{
			Status:  httpCode,
			Message: "successfully update a role",
			Data:    nil,
		},
	)
}

// @Tags 					roles
// @Summary				Delete Role By ID
// @Description 	Delete role by its id
// @Accept 				json
// @Produce 			json
// @Success				200 {object} common.BaseResponse[any]
// @Param         id path string true "Role ID"
// @Router				/roles/{id} [delete]
func (h *roleHandler) DeleteByID(c *gin.Context) {
	roleID := c.Param("id")
	parsedRoleID, err := uuid.Parse(roleID)
	if err != nil {
		errMsg := fmt.Sprintf("failed to parsed %s as roleID", roleID)
		tools.HandleLogError(err, errMsg)
		tools.HandlerSimpleError(c, http.StatusBadRequest, errMsg, err)
		return
	}

	httpCode, err := h.roleUsecase.DeleteByID(context.Background(), parsedRoleID)
	if err != nil {
		errMsg := "failed to delete role"
		tools.HandleLogError(err, errMsg)
		tools.HandlerSimpleError(c, httpCode, errMsg, err)
	}

	c.JSON(
		httpCode,
		common.BaseResponse[any]{
			Status:  httpCode,
			Message: "successfully delete role",
			Data:    nil,
		},
	)
}
