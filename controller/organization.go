package controller

import (
	"gin-demo01/entity"
	"gin-demo01/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func OrganizationController(app *gin.Engine) {
	route := app.Group("api/organization")
	{
		route.GET("", func(c *gin.Context) {
			organization := service.GetOrganizationList()
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusOK,
				"message": "Get organization successfully.",
				"data":    organization,
			})
		})

		route.POST("", func(c *gin.Context) {
			var organization entity.OrganizationEntity
			c.ShouldBind(&organization)
			service.CreateOrganization(organization)
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusOK,
				"message": "Create organization successfully.",
			})
		})
	}
}
