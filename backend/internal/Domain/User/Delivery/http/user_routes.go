package http

import (
	"github.com/gin-gonic/gin"
)

type UserRoutes struct {
	UserController *UserController
	AuthController *AuthController
}

func NewUserRoutes(userController *UserController, authController *AuthController) *UserRoutes {
	return &UserRoutes{
		UserController: userController,
		AuthController: authController,
	}
}

func (u *UserRoutes) RegisterUserRoutes(r *gin.Engine) {
	r.POST("/register", u.AuthController.Register)
	r.POST("/login", u.AuthController.Login)

	r.GET("/users", u.UserController.GetAllUsers)
	r.GET("/users/:id", u.UserController.GetUserByID)
}
