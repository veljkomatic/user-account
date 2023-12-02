package routers

import (
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"

	userdelivery "github.com/veljkomatic/user-account/pkg/user/api/delivery/rest"
	"github.com/veljkomatic/user-account/pkg/user/api/handler"
)

// InitRouter initializes a new Gin router with predefined routes and middleware
func InitRouter(userHandler handler.UserHandler) *gin.Engine {
	r := gin.New()
	// Set middleware for cors
	r.Use(cors.Default())
	// Middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// User routes
	userGroup := r.Group("/user")
	userdelivery.SetupRoutes(userGroup, userHandler)

	return r
}
