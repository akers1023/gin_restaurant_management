package routes

import (
	controller "restaurant_management/controllers/user"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.POST("users/register", controller.Register())
	incomingRoutes.POST("users/login", controller.Login())
}