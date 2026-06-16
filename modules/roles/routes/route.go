package routes

import (
	"project-root/modules/roles/providers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(rg *gin.RouterGroup, roleProvider *providers.Provider) {
	roleRoutes := rg.Group("/roles")

	roleRoutes.GET("", roleProvider.RoleHandler.GetAll)
	roleRoutes.GET("/:id", roleProvider.RoleHandler.GetByID)
	roleRoutes.POST("", roleProvider.RoleHandler.Create)
	roleRoutes.PUT("/:id", roleProvider.RoleHandler.UpdateByID)
	roleRoutes.DELETE("/:id", roleProvider.RoleHandler.DeleteByID)
}
