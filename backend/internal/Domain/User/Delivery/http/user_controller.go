package http

import (
	usecase "backend/internal/Domain/User/UseCase"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type UserController struct {
    UserService *usecase.UserService
}

func NewUserController(userService *usecase.UserService) *UserController {
    return &UserController{UserService: userService}
}

func (controller *UserController) GetUserBySubOrID(c *gin.Context) {
    userSUB := c.Param("id")

    user, err := controller.UserService.GetUserBySubOrID(userSUB)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"user": user})
}


func (controller *UserController) GetAllUsers(c *gin.Context) {
    username := c.Query("username")

    filter := bson.M{}
    if username != "" {
        filter["username"] = bson.M{"$regex": username, "$options": "i"} 
    }

    users, err := controller.UserService.GetAllUsers(filter)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"users": users})
}
