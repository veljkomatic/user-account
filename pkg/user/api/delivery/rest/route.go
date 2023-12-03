package rest

import (
	"github.com/gin-gonic/gin"
	
	"github.com/veljkomatic/user-account/pkg/user/api/handler"
)

func SetupRoutes(
	userGroup *gin.RouterGroup,
	handler handler.UserHandler,
) {
	controller := newUserController(handler)
	setupRoutes(userGroup, controller)
}

func setupRoutes(userGroup *gin.RouterGroup, controller *userController) {
	userGroup.
		GET("/:userID", controller.GetUser).
		POST("/", controller.CreateUser)
}
