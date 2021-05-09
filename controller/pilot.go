package controller

import (
	"gin-demo01/dto"
	"gin-demo01/entity"
	"gin-demo01/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func PilotController(app *gin.Engine) {
	route := app.Group("/api/pilot")
	{
		route.GET("", func(c *gin.Context) {
			var query dto.QueryPilotDTO
			var pagination dto.PaginationDTO
			query.Name = c.Query("name")
			query.OrganizationID, _ = strconv.Atoi(c.Query("organizationId"))
			pagination.Page, _ = strconv.Atoi(c.Query("page"))
			pagination.PageSize, _ = strconv.Atoi(c.Query("pageSize"))
			c.ShouldBindQuery(&query)
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusOK,
				"message": "Get pilot successfully.",
				"data":    service.GetPilotList(query, pagination),
			})
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
