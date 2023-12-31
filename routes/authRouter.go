package routes

import (
	"github.com/fredele20/golang-jwt-project/controllers"
	"github.com/gin-gonic/gin"
)

// var route controllers.ControllerService

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/signup", controllers.Signup())
	incomingRoutes.POST("users/login", controllers.Login())
	incomingRoutes.DELETE("users/logout", controllers.Logout())
	incomingRoutes.POST("users/forgot-password", controllers.ForgotPassword())
	incomingRoutes.POST("users/reset-password", controllers.ResetPassword())
}