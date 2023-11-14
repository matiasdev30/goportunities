package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")

	{
		v1.GET("opeening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "openning",
			})
		})

		v1.POST("opeening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "openning", 
			})
		})

		v1.DELETE("opeening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "openning",
			})
		})

		v1.PUT("opeening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "openning",
			})
		})


	}

}
