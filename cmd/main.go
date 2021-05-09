package main

import (
	"gin-demo01/controller"
	"gin-demo01/entity"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	entity.ConnectDB()
	// r.GET("/user", controller)
	injectController(app)
	app.Run()
}

func injectController(app *gin.Engine) {
	controller.GundamController(app)
	controller.OrganizationController(app)
	controller.PilotController(app)
}
