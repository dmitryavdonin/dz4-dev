package routes

import (
	"user-service/controllers"

	"github.com/gin-gonic/gin"
)

type UserRouteController struct {
	userController controllers.UserController
}

func NewRouteUserController(userController controllers.UserController) UserRouteController {
	return UserRouteController{userController}
}

func (uc *UserRouteController) UserRoute(rg *gin.RouterGroup) {

	router := rg.Group("user")
	router.POST("/", uc.userController.CreateUser)
	router.PUT("/:userId", uc.userController.UpdateUser)
	router.GET("/", uc.userController.FindUsers)
	router.GET("/:userId", uc.userController.FindUserById)
	router.DELETE("/:userId", uc.userController.DeleteUser)
}
