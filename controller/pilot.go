package controller

import (
	"gin-demo01/dto"
	"gin-demo01/entity"
	"gin-demo01/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PilotController(app *gin.Engine) {
	route := app.Group("/api/pilot")
	{
		route.GET("", func(c *gin.Context) {
			var query dto.QueryPilotDTO
			if err := c.ShouldBindQuery(&query); err == nil {
				c.JSON(http.StatusOK, gin.H{
					"code":    http.StatusOK,
					"message": "Get pilot successfully.",
					"data":    service.GetPilotList(query),
				})
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
		})

		route.POST("", func(c *gin.Context) {
			var pilot entity.PilotEntity
			c.ShouldBind(&pilot)
			service.CreatePilot(pilot)
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusOK,
				"message": "Create pilot successfully.",
			})
		})
	}
}
