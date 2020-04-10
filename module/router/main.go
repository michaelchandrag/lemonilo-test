package router

import (
	"github.com/gin-gonic/gin"

	controller "github.com/michaelchandrag/lemonilo-test/module/controller"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/user_profiles", controller.Find)
	r.POST("/user_profile", controller.Add)
	r.PUT("/user_profile/:user_id", controller.Update)
	r.DELETE("/user_profile/:user_id", controller.Delete)

	r.POST("/login", controller.Login)

	return r
}