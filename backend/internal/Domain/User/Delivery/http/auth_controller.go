package http

import (
	domain "backend/internal/Domain/User/Domain"
	ports "backend/internal/Domain/User/Ports"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	UserPort ports.UserPort
}

func NewAuthController(userPort ports.UserPort) *AuthController {
	return &AuthController{
		UserPort: userPort,
	}
}

func (c *AuthController) Register(ctx *gin.Context) {
	var user domain.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if user.Email == "" || user.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Email and password are required"})
		return
	}

	err := c.UserPort.Register(&user)
	if err != nil {
		if err.Error() == "email already registered" {
			ctx.JSON(http.StatusConflict, gin.H{"error": "Email already in use"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func (c *AuthController) Login(ctx *gin.Context) {
	var creds domain.User
	if err := ctx.ShouldBindJSON(&creds); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := c.UserPort.Login(creds.Email, creds.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
