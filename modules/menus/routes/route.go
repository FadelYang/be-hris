package routes

import (
	"project-root/modules/menus/providers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(rg *gin.RouterGroup, roleProvider *providers.Provider) {
	roleRoutes := rg.Group("/menus")

	roleRoutes.GET("", roleProvider.MenuHandler.GetAll)
	roleRoutes.GET("/:id", roleProvider.MenuHandler.GetByID)
	roleRoutes.POST("", roleProvider.MenuHandler.Create)
	roleRoutes.PUT("/:id", roleProvider.MenuHandler.UpdateByID)
	roleRoutes.DELETE("/:id", roleProvider.MenuHandler.DeleteByID)
}
