package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/veljkomatic/user-account/pkg/user/api"
	"github.com/veljkomatic/user-account/pkg/user/api/handler"
	"net/http"
)

// rest controller for user
type userController struct {
	handler handler.UserHandler
}

func newUserController(handler handler.UserHandler) *userController {
	return &userController{
		handler: handler,
	}
}

func (uc *userController) GetUser(c *gin.Context) {
	var req handler.GetUserRequest
	err := c.ShouldBindUri(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx := c.Request.Context()
	apiCtx := api.NewAPIContext(ctx)

	userResponse, err := uc.handler.GetUser(apiCtx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if userResponse == nil {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	c.JSON(userResponse.StatusCode(), userResponse)
}

func (uc *userController) CreateUser(c *gin.Context) {
	var req handler.CreateUserRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := req.ValidateRequest(); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx := c.Request.Context()
	apiCtx := api.NewAPIContext(ctx)

	userResponse, err := uc.handler.CreateUser(apiCtx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(userResponse.StatusCode(), userResponse)
}
