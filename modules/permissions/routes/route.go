package routes

import (
	"project-root/modules/permissions/providers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(rg *gin.RouterGroup, roleProvider *providers.Provider) {
	roleRoutes := rg.Group("/permissions")

	roleRoutes.GET("", roleProvider.PermissionHandler.GetAll)
	roleRoutes.GET("/:id", roleProvider.PermissionHandler.GetByID)
	roleRoutes.POST("", roleProvider.PermissionHandler.Create)
	roleRoutes.PUT("/:id", roleProvider.PermissionHandler.UpdateByID)
	roleRoutes.DELETE("/:id", roleProvider.PermissionHandler.DeleteByID)
}
