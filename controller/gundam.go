package controller

import (
	"gin-demo01/dto"
	"gin-demo01/entity"
	"gin-demo01/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GundamController(app *gin.Engine) {
	route := app.Group("/api/gundam")
	{
		route.GET("", func(c *gin.Context) {
			var query dto.QueryGundamDTO
			var pagination dto.PaginationDTO
			query.Name = c.Query("name")
			query.Code = c.Query("code")
			pagination.Page, _ = strconv.Atoi(c.Query("page"))
			pagination.PageSize, _ = strconv.Atoi(c.Query("pageSize"))
			gundam := service.GetGundamList(query, pagination)
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusOK,
				"message": "Get gundam successfully.",
				"data":    gundam,
			})
		})
		route.POST("", func(c *gin.Context) {
			var gundam entity.GundamEntity
			c.ShouldBind(&gundam)
			service.CreateGundam(gundam)
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusOK,
				"message": "Create gundam succesfully.",
			})
		})
		route.POST("/batch", func(c *gin.Context) {
			var gundam []entity.GundamEntity
			c.ShouldBind(&gundam)
			service.BatchCreateGundam(gundam)
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusOK,
				"message": "Batch create gundam succesfully.",
			})
		})
		route.PUT("/:id", func(c *gin.Context) {
			var gundam entity.GundamEntity
			id, _ := strconv.Atoi(c.Param("id"))
			c.ShouldBind(&gundam)
			gundam.ID = uint(id)
			service.UpdateGundam(gundam)
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusOK,
				"message": "Update gundam succesfully.",
			})
		})
		route.DELETE("/:id", func(c *gin.Context) {
			id, _ := strconv.Atoi(c.Param("id"))
			service.DeleteGundam(uint(id))
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusOK,
				"message": "Delete gundam succesfully.",
			})
		})
	}
}
